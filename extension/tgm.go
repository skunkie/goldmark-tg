package extension

import (
	"github.com/yuin/goldmark"
)

type tgm struct{}

// TGM is an extension that provides Telegram markdown functionalities.
var TGM = &tgm{}

func (e *tgm) Extend(m goldmark.Markdown) {
	Emphasis.Extend(m)
}
