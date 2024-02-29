package extension

import (
	"github.com/yuin/goldmark"
)

type tgmV2 struct{}

// TGMV2 is an extension that provides Telegram markdownV2 functionalities.
var TGMV2 = &tgmV2{}

func (e *tgmV2) Extend(m goldmark.Markdown) {
	EmphasisV2.Extend(m)
	Spoiler.Extend(m)
	Strikethrough.Extend(m)
}
