package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

// A Emphasis struct represents a emphasis of Telegram Markdown text.
type Emphasis struct {
	gast.BaseInline

	Delimiter byte
	Level     int
}

// Dump implements Node.Dump.
func (n *Emphasis) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindEmphasis is a NodeKind of the Emphasis node.
var KindEmphasis = gast.NewNodeKind("Emphasis")

// Kind implements Node.Kind.
func (n *Emphasis) Kind() gast.NodeKind {
	return KindEmphasis
}

// NewEmphasis returns a new Emphasis node.
func NewEmphasis(delimiter byte, level int) *Emphasis {
	return &Emphasis{Delimiter: delimiter, Level: level}
}
