// Deprecated: Use the `output` package instead.
package color

import (
	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed)
	Green  = color.New(color.FgGreen)
	Yellow = color.New(color.FgYellow)

	// Deprecated: Use the `output` package instead.
	Error = Red.Add(color.Bold)
	// Deprecated: Use the `output` package instead.
	Warn = Yellow.Add(color.Underline)
)
