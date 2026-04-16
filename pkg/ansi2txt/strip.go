package ansi2txt

import (
	"strings"
)

// Strip removes ANSI escape sequences from s.
func Strip(s string) string {
	var buf strings.Builder
	buf.Grow(len(s))
	w := NewWriter(&buf)
	_, _ = w.Write([]byte(s))
	return buf.String()
}
