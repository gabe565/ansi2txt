package ansi2txt

import (
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{"no escape", "abc", "abc", require.NoError},
		{"blue", color.BlueString("blue"), "blue", require.NoError},
		{"formatted", color.New(color.Bold, color.Italic).Sprint("formatted"), "formatted", require.NoError},
		{"move cursor", "\x1b[10;10Hmove", "move", require.NoError},
		{
			"complex",
			"foo\x1b[1mbar\x1b[0m\n" +
				"\x1b[?2004l" +
				"\x1b]0;title1\x1b\\" +
				"\x1b]4;16;rgb:0/0/0\x1b\\" +
				"\x1b]0;title2\a" +
				"\n",
			"foobar\n\n",
			require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			w := NewWriter(&buf)
			got, err := w.Write([]byte(tt.input))
			tt.wantErr(t, err)
			assert.Equal(t, len(tt.input), got)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
