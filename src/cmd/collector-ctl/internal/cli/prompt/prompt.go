package prompt

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/internal/repositories/database"
)

type Prompt struct {
	*database.Database
}

func (p *Prompt) TryAgain(warning string) bool {
	color.Warn.Println(warning)

	return p.Confirm("Would you like to try again?", true)
}

func (p *Prompt) Confirm(message string, defaultValue bool) (retry bool) {
	question := &survey.Confirm{
		Message: message,
		Default: defaultValue,
		Help:    "If you don't know what to do, just press enter.",
	}

	if err := survey.AskOne(question, &retry); err != nil {
		color.Error.Printf("Failed to confirm with error: %s\n", err.Error())
		return false
	}
	return
}

var (
	ErrResultNotFound = errors.New("result not found")
)
