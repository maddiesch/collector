package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var (
	EnDisplay = display.English.Languages()
)

func DisplayNameForTagString(s string) string {
	switch s {
	case "ph":
		return "Phyrexian"
	default:
		tag, err := language.Parse(s)
		if err != nil {
			return s
		}
		return DisplayNameForTag(tag)
	}
}

func DisplayNameForTag(tag language.Tag) string {
	return EnDisplay.Name(tag)
}
