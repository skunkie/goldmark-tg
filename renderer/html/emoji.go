// Package html implements renderer that outputs HTMLs.
package html

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

const emojiPrefix = "tg://emoji?id="

type EmojiHTMLRenderer struct {
	html.Config
}

// NewEmojiHTMLRenderer returns a new EmojiHTMLRenderer.
func NewEmojiHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &EmojiHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *EmojiHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(gast.KindImage, r.renderEmoji)
}

func (r *EmojiHTMLRenderer) renderEmoji(w util.BufWriter, source []byte, node gast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*gast.Image)
	prefix := []byte(emojiPrefix)
	if bytes.HasPrefix(n.Destination, prefix) {
		if entering {
			_, _ = w.WriteString(`<tg-emoji emoji-id="`)
			id := bytes.TrimPrefix(n.Destination, prefix)
			_, _ = w.WriteString(string(util.EscapeHTML(id)))
			_, _ = w.WriteString(`">`)
		} else {
			_, _ = w.WriteString("</tg-emoji>")
		}
	}

	return gast.WalkContinue, nil
}
