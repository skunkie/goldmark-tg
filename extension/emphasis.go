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

type Config struct {
	Writer html.Writer
	Level  int
}

func NewConfig() Config {
	return Config{
		Writer: html.DefaultWriter,
	}
}

type withLevel struct {
	Level int
}

func (o *withLevel) SetHTMLOption(c *Config) {
	c.Level = o.Level
}

func WithLevel(level int) interface {
	Option
} {
	return &withLevel{Level: level}
}

type Option interface {
	SetHTMLOption(*Config)
}

type emphasisDelimiterProcessor struct {
	Delimiter byte
}

func (p *emphasisDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == p.Delimiter
}

func (p *emphasisDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *emphasisDelimiterProcessor) OnMatch(consumes int) gast.Node {
	return ast.NewEmphasis(p.Delimiter, consumes)
}

type emphasisParser struct{}

var defaultEmphasisParser = &emphasisParser{}

// NewEmphasisParser return a new InlineParser that parses
// emphasis expressions.
func NewEmphasisParser() parser.InlineParser {
	return defaultEmphasisParser
}

func (s *emphasisParser) Trigger() []byte {
	return []byte{'*', '_'}
}

func (s *emphasisParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 1, &emphasisDelimiterProcessor{Delimiter: '*'})
	if node == nil {
		node = parser.ScanDelimiter(line, before, 1, &emphasisDelimiterProcessor{Delimiter: '_'})
		if node == nil {
			return nil
		}
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *emphasisParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

// EmphasisHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Emphasis nodes.
type EmphasisHTMLRenderer struct {
	Config
}

// NewEmphasisHTMLRenderer returns a new EmphasisHTMLRenderer.
func NewEmphasisHTMLRenderer(opts ...Option) renderer.NodeRenderer {
	r := &EmphasisHTMLRenderer{
		Config: NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *EmphasisHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindEmphasis, r.renderEmphasis)
}

func (r *EmphasisHTMLRenderer) renderEmphasis(
	w util.BufWriter, source []byte, node gast.Node, entering bool,
) (gast.WalkStatus, error) {
	var (
		n   = node.(*ast.Emphasis)
		tag string
	)
	switch {
	case n.Level == 1 && n.Delimiter == '*':
		tag = "b"
	case n.Level == 1 && n.Delimiter == '_':
		tag = "i"
	case r.Level == 2 && n.Level == 2 && n.Delimiter == '_':
		tag = "u"
	default:
		return gast.WalkContinue, nil
	}

	if entering {
		_, _ = w.WriteString("<" + tag + ">")
	} else {
		_, _ = w.WriteString("</" + tag + ">")
	}
	return gast.WalkContinue, nil
}

type emphasis struct {
	option Option
}

var (
	// Emphasis is an extension that allow you to use emphasis expressions like '*bold*', '_italic_'.
	Emphasis = &emphasis{WithLevel(1)}
	// EmphasisV2 is an extension that allow you to use emphasis expressions like '*bold*', '_italic_' or '__undeline__'.
	EmphasisV2 = &emphasis{WithLevel(2)}
)

func (e *emphasis) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewEmphasisParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewEmphasisHTMLRenderer(e.option), 500),
	))
}
