package tg_test

import (
	"os"
	"strings"
	"testing"

	tg "github.com/skunkie/goldmark-tg"
	"github.com/skunkie/goldmark-tg/extension"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
)

func readfile(t *testing.T, path string) []byte {
	t.Helper()

	source, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	return source
}

func TestMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name   string
		md     goldmark.Markdown
		source []byte
		want   string
	}{
		{
			name:   "test 01",
			md:     goldmark.New(tg.Markdown()...),
			source: readfile(t, "_test/markdown.md"),
			want:   "<b>bold text</b>\n<i>italic text</i>\n<a href=\"http://www.example.com/\">inline URL</a>\n<a href=\"tg://user?id=123456789\">inline mention of a user</a>\n<code>inline fixed-width code</code><pre><code>pre-formatted fixed-width code block\n</code></pre>\n<pre><code class=\"language-python\">pre-formatted fixed-width code block written in the Python programming language\n</code></pre>\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := tt.md.Convert(tt.source, &buf)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestMarkdownV2ToHTML(t *testing.T) {
	tests := []struct {
		name   string
		md     goldmark.Markdown
		source []byte
		want   string
	}{
		{
			name:   "test 01",
			md:     goldmark.New(tg.MarkdownV2()...),
			source: readfile(t, "_test/markdownv2.md"),
			want:   "<b>bold *text</b>\n<i>italic *text</i>\n<u>underline</u>\n<s>strikethrough</s>\n<tg-spoiler>spoiler</tg-spoiler>\n<b>bold <i>italic bold <s>italic bold strikethrough <tg-spoiler>italic bold strikethrough spoiler</tg-spoiler></s> <u>underline italic bold</u></i> bold</b>\n<a href=\"http://www.example.com/\">inline URL</a>\n<a href=\"tg://user?id=123456789\">inline mention of a user</a>\n<tg-emoji emoji-id=\"5368324170671202286\">👍</tg-emoji>\n<code>inline fixed-width code</code><pre><code>pre-formatted fixed-width code block\n</code></pre>\n<pre><code class=\"language-python\">pre-formatted fixed-width code block written in the Python programming language\n</code></pre>\n<blockquote>\nBlock quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>\n",
		},
		{
			name:   "test 02",
			md:     goldmark.New(append(tg.MarkdownV2(), goldmark.WithExtensions(extension.Template))...),
			source: readfile(t, "_test/markdownv2_template.md"),
			want:   "<b>{{ if eq .Status \"firing\" }}2{{ else }}0{{ end }}</b>\n<i>T1</i> <u>{{ __T2__\n||\"T3\"|| }}</u>\n}} <tg-spoiler>&quot;T4&quot;</tg-spoiler> {{",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := tt.md.Convert(tt.source, &buf)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
