package ast

import (
	gast "github.com/yuin/goldmark/ast"
	textm "github.com/yuin/goldmark/text"
)

// A Template struct represents a Go template like {{ if eq .Status "firing" }}.
type Template struct {
	gast.BaseInline

	Segment textm.Segment
}

// Dump implements Node.Dump.
func (n *Template) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindTemplate is a NodeKind of the Template node.
var KindTemplate = gast.NewNodeKind("Template")

// Kind implements Node.Kind.
func (n *Template) Kind() gast.NodeKind {
	return KindTemplate
}

// NewTemplate returns a new Template node.
func NewTemplate() *Template {
	return &Template{}
}
