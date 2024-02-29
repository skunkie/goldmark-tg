package main

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark"

	tg "github.com/skunkie/goldmark-tg"
	"github.com/skunkie/goldmark-tg/extension"
)

var template = []byte(`*{{ if eq .Status "firing" }}2{{ else }}0{{ end }}*`)

func main() {
	var buf strings.Builder
	md := goldmark.New(append(tg.MarkdownV2(), goldmark.WithExtensions(extension.Template))...)
	if err := md.Convert(template, &buf); err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String()) //nolint:forbidigo
}
