package extension

import (
	"regexp"

	"github.com/skunkie/goldmark-tg/extension/ast"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type templateParser struct{}

var defaultTemplateParser = &templateParser{}

// NewTemplateParser return a new InlineParser that parses
// template expressions.
func NewTemplateParser() parser.InlineParser {
	return defaultTemplateParser
}

func (s *templateParser) Trigger() []byte {
	return []byte{'{'}
}

var templateRegexp = regexp.MustCompile("^{{[^{}]+}}")

func (s *templateParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	_, segment := block.Position()
	if !block.Match(templateRegexp) {
		return nil
	}
	_, current := block.Position()
	segment = segment.WithStop(current.Start)
	node := ast.NewTemplate()
	node.Segment = segment
	return node
}

func (s *templateParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

// TemplateHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Template nodes.
type TemplateHTMLRenderer struct {
	html.Config
}

// NewTemplateHTMLRenderer returns a new TemplateHTMLRenderer.
func NewTemplateHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &TemplateHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *TemplateHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindTemplate, r.renderTemplate)
}

func (r *TemplateHTMLRenderer) renderTemplate(
	w util.BufWriter, source []byte, node gast.Node, entering bool,
) (gast.WalkStatus, error) {
	n := node.(*ast.Template)
	if entering {
		_, _ = w.Write(n.Segment.Value(source))
	}
	return gast.WalkContinue, nil
}

type template struct{}

// Template is an extension that allow you to use template expression like '{{.Count}}'.
var Template = &template{}

func (e *template) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewTemplateParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTemplateHTMLRenderer(), 500),
	))
}
