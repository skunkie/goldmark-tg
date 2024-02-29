goldmark-tg
=========================

goldmark-tg is a collection of extensions for the [goldmark](http://github.com/yuin/goldmark) 
that adds Telegram Markdown and MarkdownV2 functionalities.

### Examples

Markdown:

```go
package main

import (
	"github.com/yuin/goldmark"

	tg "github.com/skunkie/goldmark-tg"
)

func main() {
	md := goldmark.New(tg.Markdown()...)
	...
}
```

MarkdownV2:

```go
package main

import (
	"github.com/yuin/goldmark"

	tg "github.com/skunkie/goldmark-tg"
)

func main() {
	md := goldmark.New(tg.MarkdownV2()...)
	...
}
```

MarkdownV2 with Go Template extension:

```go
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
	fmt.Printf("%q\n", buf.String())
}
```