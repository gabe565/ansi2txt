package ansi2txt

import (
	"io"
	"slices"
)

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// A Writer writes data with ANSI escape sequences removed.
type Writer struct {
	w     io.Writer
	buf   []byte
	state state
}

type state uint8

const (
	stateDefault state = iota
	stateEscaped
	stateInCSI
	stateStartOSC
	stateInOSC
	stateIgnoreNext
)

const (
	escape byte = 0x1B
	bell   byte = 0x7
)

// Write writes data to w with ANSI escape sequences removed.
func (w *Writer) Write(p []byte) (int, error) {
	w.buf = slices.Grow(w.buf[:0], len(p))

	for _, b := range p {
		switch w.state {
		case stateDefault:
			switch b {
			case escape:
				w.state = stateEscaped
			case bell:
			default:
				w.buf = append(w.buf, b)
			}
		case stateEscaped:
			switch b {
			case '[':
				w.state = stateInCSI
			case ']':
				w.state = stateStartOSC
			case '%', '(', ')', '0', '3', '5', '6', '#':
				w.state = stateIgnoreNext
			case bell, 'A', 'B', 'C', 'D', 'E', 'H', 'I', 'J', 'K', 'M', 'N', 'O', 'S', 'T', 'Z', 'c', 's', 'u', '1', '2', '7', '8', '<', '=', '>':
				w.state = stateDefault
			default:
				w.buf = append(w.buf, escape, b)
				w.state = stateDefault
			}
		case stateInCSI:
			switch b {
			case ';', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				w.state = stateDefault
			}
		case stateStartOSC:
			switch b {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				w.state = stateInOSC
			default:
				w.state = stateDefault
			}
		case stateInOSC:
			switch b {
			case bell:
				w.state = stateDefault
			case escape:
				w.state = stateIgnoreNext
			}
		case stateIgnoreNext:
			w.state = stateDefault
		}
	}

	n, err := w.w.Write(w.buf)
	switch {
	case err != nil:
		return n, err
	case n < len(w.buf):
		return n, io.ErrShortWrite
	default:
		return len(p), nil
	}
}

// Reset clears the internal state.
func (w *Writer) Reset() {
	w.buf = nil
	w.state = stateDefault
}
