// Package html implements renderer that outputs HTMLs.
package html

import (
	"github.com/yuin/goldmark/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type ParagraphHTMLRenderer struct {
	html.Config
}

// NewParagraphHTMLRenderer returns a new ParagraphHTMLRenderer.
func NewParagraphHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &ParagraphHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *ParagraphHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(gast.KindParagraph, r.renderParagraph)
}

func (r *ParagraphHTMLRenderer) renderParagraph(w util.BufWriter, source []byte, node gast.Node, entering bool) (ast.WalkStatus, error) {
	return gast.WalkContinue, nil
}
