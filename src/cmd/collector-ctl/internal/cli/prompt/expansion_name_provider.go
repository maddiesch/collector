package prompt

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/ports"
)

func (p *Prompt) ProvideExpansionName(ctx context.Context, cardName string) (string, error) {
	question := &survey.Input{
		Message: "Default expansion name:",
		Help:    "This will be used for the initial card name search. Usefull if you're entering multiple cards from the same expansion.",
		Suggest: func(toComplete string) []string {
			names, _ := p.ExpansionNameSearchPrefix(ctx, toComplete, cardName)
			return names
		},
	}

	var inputName string

Retry:
	if err := survey.AskOne(question, &inputName); err != nil {
		return "", err
	}
	if inputName == "" {
		return "", nil
	}

	if exists, err := p.ExpansionNameExists(ctx, inputName); err != nil {
		return "", err
	} else if !exists {
		if p.TryAgain("Entered expansion name does not exist") {
			goto Retry
		}
		return "", ErrResultNotFound
	}

	return inputName, nil
}

var _ ports.ExpansionNameProvider = (*Prompt)(nil)
