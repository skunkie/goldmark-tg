package extension

import (
	"github.com/skunkie/goldmark-tg/extension/ast"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type spoilerDelimiterProcessor struct{}

func (p *spoilerDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '|'
}

func (p *spoilerDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *spoilerDelimiterProcessor) OnMatch(consumes int) gast.Node {
	return ast.NewSpoiler()
}

var defaultSpoilerDelimiterProcessor = &spoilerDelimiterProcessor{}

type spoilerParser struct{}

var defaultSpoilerParser = &spoilerParser{}

// NewSpoilerParser return a new InlineParser that parses
// spoiler expressions.
func NewSpoilerParser() parser.InlineParser {
	return defaultSpoilerParser
}

func (s *spoilerParser) Trigger() []byte {
	return []byte{'|'}
}

func (s *spoilerParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultSpoilerDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *spoilerParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

// SpoilerHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Spoiler nodes.
type SpoilerHTMLRenderer struct {
	html.Config
}

// NewSpoilerHTMLRenderer returns a new SpoilerHTMLRenderer.
func NewSpoilerHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &SpoilerHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *SpoilerHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindSpoiler, r.renderSpoiler)
}

func (r *SpoilerHTMLRenderer) renderSpoiler(
	w util.BufWriter, source []byte, n gast.Node, entering bool,
) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<tg-spoiler>")
	} else {
		_, _ = w.WriteString("</tg-spoiler>")
	}
	return gast.WalkContinue, nil
}

type spoiler struct{}

// Spoiler is an extension that allow you to use spoiler expression like '||spoiler||'.
var Spoiler = &spoiler{}

func (e *spoiler) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewSpoilerParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewSpoilerHTMLRenderer(), 500),
	))
}
