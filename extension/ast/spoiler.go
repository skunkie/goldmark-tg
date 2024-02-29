package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

// A Spoiler struct represents a spoiler of Telegram MarkdownV2 text.
type Spoiler struct {
	gast.BaseInline
}

// Dump implements Node.Dump.
func (n *Spoiler) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindSpoiler is a NodeKind of the Spoiler node.
var KindSpoiler = gast.NewNodeKind("Spoiler")

// Kind implements Node.Kind.
func (n *Spoiler) Kind() gast.NodeKind {
	return KindSpoiler
}

// NewSpoiler returns a new Spoiler node.
func NewSpoiler() *Spoiler {
	return &Spoiler{}
}
