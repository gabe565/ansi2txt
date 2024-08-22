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

// Write writes data to w with ANSI escape sequences removed.
func (w *Writer) Write(p []byte) (int, error) {
	for i, b := range p {
		switch w.state {
		case stateNone:
			switch b {
			case '\x1b':
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
				if b == '\a' { // Bell
					w.state = stateNone
				} else if b == '\'' {
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
