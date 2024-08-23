package ansi2txt

import (
	"bufio"
	"io"
)

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: bufio.NewWriter(w)}
}

// A Writer writes data with ANSI escape sequences removed.
type Writer struct {
	w     *bufio.Writer
	state state
}

type state uint8

const (
	stateNone state = iota
	stateEscape
	stateCSI
	stateOSC
	stateIgnore
)

const (
	escape = 0x1B
	bell   = 0x7
)

// Write writes data to w with ANSI escape sequences removed.
func (w *Writer) Write(p []byte) (int, error) {
	for i, b := range p {
		switch w.state {
		case stateNone:
			switch b {
			case escape:
				w.state = stateEscape
			default:
				if err := w.w.WriteByte(b); err != nil {
					return i, err
				}
			}
		case stateEscape:
			switch b {
			case '[':
				w.state = stateCSI
			case ']':
				w.state = stateOSC
			case '%', '(', ')':
				w.state = stateIgnore
			}
		case stateCSI:
			if b != ';' && (b < '0' || b > '9') && b != '?' {
				w.state = stateNone
			}
		case stateOSC:
			if b <= '9' {
				switch b {
				case bell:
					w.state = stateNone
				case escape:
					w.state = stateIgnore
				}
			}
		case stateIgnore:
			w.state = stateNone
		}
	}

	return len(p), w.w.Flush()
}

// Reset clears the internal state.
func (w *Writer) Reset() {
	w.state = stateNone
}
