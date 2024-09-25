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
	stateNone state = iota
	stateEscape
	stateCSI
	stateOSCFirst
	stateOSC
	stateIgnore
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
		case stateNone:
			switch b {
			case escape:
				w.state = stateEscape
			case bell:
			default:
				w.buf = append(w.buf, b)
			}
		case stateEscape:
			switch b {
			case '[':
				w.state = stateCSI
			case ']':
				w.state = stateOSCFirst
			case '%', '(', ')', '0', '3', '5', '6', '#':
				w.state = stateIgnore
			case bell, 'A', 'B', 'C', 'D', 'E', 'H', 'I', 'J', 'K', 'M', 'N', 'O', 'S', 'T', 'Z', 'c', 's', 'u', '1', '2', '7', '8', '<', '=', '>':
				w.state = stateNone
			default:
				w.buf = append(w.buf, b)
				w.state = stateNone
			}
		case stateCSI:
			switch b {
			case ';', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				w.state = stateNone
			}
		case stateOSCFirst:
			switch b {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				w.state = stateOSC
			default:
				w.state = stateNone
			}
		case stateOSC:
			switch b {
			case bell:
				w.state = stateNone
			case escape:
				w.state = stateIgnore
			}
		case stateIgnore:
			w.state = stateNone
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
	w.state = stateNone
}
