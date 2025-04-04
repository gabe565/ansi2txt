package ansi2txt

import (
	"regexp"
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
		{"only escape", "a\x1bbc", "a\x1bbc", require.NoError},
		{"character set", "\x1b%Gabc", "abc", require.NoError},
		{"hide cursor", "\x1b[?25lhidden\x1b[?25h", "hidden", require.NoError},
		{"invalid escape sequence", "a\x1b[999mbc", "abc", require.NoError},
		{
			"sequences mixed with text",
			"normal \x1b[1mbold\x1b[0m normal \x1b[4munderlined\x1b[0m",
			"normal bold normal underlined",
			require.NoError,
		},
		{`title (ST ESC \)`, "\x1b]2;title\x1b\"abc", "abc", require.NoError},
		{"title (ST BEL)", "\x1b]0;title\x07abc", "abc", require.NoError},
		{
			`hyperlink (ST ESC \)`,
			"\x1b]8;;https://example.com\x1b\\example.com\x1b]8;;\x1b\\",
			"example.com",
			require.NoError,
		},
		{
			"hyperlink (ST BEL)",
			"\x1b]8;;https://example.com\x07example.com\x1b]8;;\x07",
			"example.com",
			require.NoError,
		},
		{"ls output", "\x1b[00;38;5;244m\x1b[m\x1b[00;38;5;33mfoo\x1b[0m", "foo", require.NoError},
		{"clear tabs", "foo\x1b[0gbar", "foobar", require.NoError},
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
			buf.Grow(len(tt.input))
			w := NewWriter(&buf)
			got, err := w.Write([]byte(tt.input))
			tt.wantErr(t, err)
			assert.Equal(t, len(tt.input), got)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

const consumptionCharacters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+1234567890-=[]{};':"./>?,<\|`

// From https://github.com/chalk/ansi-regex/blob/f338e1814144efb950276aac84135ff86b72dc8e/test.js#L64-L97
func TestWriter_Write_AnsiCodes(t *testing.T) {
	re, err := regexp.Compile(`\d$`) //nolint:gocritic
	require.NoError(t, err)

	for codeSetName, codeSet := range ansiCodes() {
		for code, info := range codeSet {
			t.Run(codeSetName+" "+info, func(t *testing.T) {
				if re.MatchString(code) {
					t.SkipNow()
				}

				input := "te\x1b" + code + "st"
				var buf strings.Builder
				buf.Grow(len(input))
				w := NewWriter(&buf)
				got, err := w.Write([]byte(input))
				require.NoError(t, err)
				assert.Equal(t, len(input), got)
				assert.Equal(t, "test", buf.String(), code)
			})

			t.Run(codeSetName+" "+info+" does not overconsume", func(t *testing.T) {
				if re.MatchString(code) {
					t.SkipNow()
				}

				for _, r := range consumptionCharacters {
					input := "\x1b" + code + string(r)
					var buf strings.Builder
					buf.Grow(len(input))
					w := NewWriter(&buf)
					got, err := w.Write([]byte(input))
					require.NoError(t, err)
					assert.Equal(t, len(input), got)
					assert.Equal(t, string(r), buf.String(), code)
				}
			})
		}
	}
}
