package styles

import (
	"github.com/alecthomas/chroma"
)

// VS Code style.
var DarkPlus = Register(chroma.MustNewStyle("dark-plus", chroma.StyleEntries{
	chroma.Comment:           "#303030",
	chroma.CommentPreproc:    "#608b4e",
	chroma.Keyword:           "#569cd6",
	chroma.OperatorWord:      "#c586c0",
	chroma.KeywordType:       "#4ec9b0",
	chroma.NameClass:         "#4ec9b0",
	chroma.LiteralString:     "#ce9178",
	chroma.GenericHeading:    "bold",
	chroma.GenericSubheading: "bold",
	chroma.GenericEmph:       "italic",
	chroma.GenericStrong:     "bold",
	chroma.GenericPrompt:     "bold",
	chroma.Error:             "border:#d16969",
	chroma.Background:        " bg:#ffffff",
}))
