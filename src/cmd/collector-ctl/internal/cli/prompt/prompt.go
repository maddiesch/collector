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

	question := &survey.Confirm{
		Message: "Would you like to try again?",
		Default: true,
		Help:    "If you don't know what to do, just press enter.",
	}

	var retry bool

	if err := survey.AskOne(question, &retry); err != nil {
		color.Error.Printf("Failed to ask for retry with error: %s\n", err.Error())
		return false
	}

	return retry
}

var (
	ErrResultNotFound = errors.New("result not found")
)
