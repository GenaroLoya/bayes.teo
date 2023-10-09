package fonttypes

import (
	"github.com/fatih/color"
)

// FontTypes is a function that prints the font types.
var Title = color.New(color.BgHiMagenta, color.Bold, color.Italic, color.Underline, color.BgHiBlue)
var SubTitle = color.New(color.FgHiCyan, color.Italic)
var Text = color.New(color.FgHiWhite)
