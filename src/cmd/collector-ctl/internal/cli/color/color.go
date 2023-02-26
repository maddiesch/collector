package color

import (
	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed)
	Green  = color.New(color.FgGreen)
	Yellow = color.New(color.FgYellow)

	Error = Red.Add(color.Bold)
	Warn  = Yellow.Add(color.Underline)
)
