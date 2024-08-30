package ansi2txt

import (
	"strings"
	"testing"

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
		{"color", "\x1b[34mblue", "blue", require.NoError},
		{"multiple", "\x1b[1mbold\x1b[0m and \x1b[4munderlined\x1b[0m", "bold and underlined", require.NoError},
		{"nested", "\x1b[31;1mred bold\x1b[0m normal", "red bold normal", require.NoError},
		{"clear screen", "\x1b[2Jclear", "clear", require.NoError},
		{"move cursor", "\x1b[10;10Hmove", "move", require.NoError},
		{"bell", "a\x07bc", "abc", require.NoError},
		{"only escape", "a\x1bbc", "abc", require.NoError},
		{"title", "\x1b]0;title\x07abc", "abc", require.NoError},
		{"character set", "\x1b%Gabc", "abc", require.NoError},
		{"hide cursor", "\x1b[?25lhidden\x1b[?25h", "hidden", require.NoError},
		{"invalid escape sequence", "a\x1b[999mbc", "abc", require.NoError},
		{"title terminated by escape", "\x1b]2;title\x1b\"abc", "abc", require.NoError},
		{"sequences mixed with text", "normal \x1b[1mbold\x1b[0m normal \x1b[4munderlined\x1b[0m", "normal bold normal underlined", require.NoError},
		{"unclosed sequence", "text \x1b[31 no close", "text no close", require.NoError},
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
