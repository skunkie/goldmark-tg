package tg

import (
	"github.com/skunkie/goldmark-tg/extension"
	"github.com/skunkie/goldmark-tg/renderer/html"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func blockParsers() []util.PrioritizedValue {
	return []util.PrioritizedValue{
		util.Prioritized(parser.NewCodeBlockParser(), 500),
		util.Prioritized(parser.NewFencedCodeBlockParser(), 700),
		util.Prioritized(parser.NewBlockquoteParser(), 800),
		util.Prioritized(parser.NewParagraphParser(), 1000),
	}
}

func inlineParsers() []util.PrioritizedValue {
	return []util.PrioritizedValue{
		util.Prioritized(parser.NewCodeSpanParser(), 100),
		util.Prioritized(parser.NewLinkParser(), 200),
	}
}

func tgParser() parser.Parser {
	return parser.NewParser(
		parser.WithBlockParsers(blockParsers()...),
		parser.WithInlineParsers(inlineParsers()...),
	)
}

// A collection of goldmark options to parse and render Telegram Markdown.
func Markdown() []goldmark.Option {
	return []goldmark.Option{
		goldmark.WithExtensions(extension.TGM),
		goldmark.WithParser(tgParser()),
		goldmark.WithRendererOptions(
			renderer.WithNodeRenderers(
				util.PrioritizedValue{Value: html.NewParagraphHTMLRenderer(), Priority: 500},
			),
		),
	}
}

// A collection of goldmark options to parse and render Telegram MarkdownV2.
func MarkdownV2() []goldmark.Option {
	return []goldmark.Option{
		goldmark.WithExtensions(extension.TGMV2),
		goldmark.WithParser(tgParser()),
		goldmark.WithRendererOptions(
			renderer.WithNodeRenderers(
				util.PrioritizedValue{Value: html.NewEmojiHTMLRenderer(), Priority: 500},
				util.PrioritizedValue{Value: html.NewParagraphHTMLRenderer(), Priority: 500},
			),
		),
	}
}
