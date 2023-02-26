package prompt

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/samber/lo"
)

func (p *Prompt) ProvideDefaultExpansionName(ctx context.Context) (string, error) {
	question := &survey.Input{
		Message: "Default expansion name:",
		Help:    "This will be used for the initial card name search. Usefull if you're entering multiple cards from the same expansion.",
		Suggest: func(toComplete string) []string {
			names, _ := p.ExpansionNameSearchPrefix(ctx, toComplete, "")
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

func (p *Prompt) ProvideExpansionName(ctx context.Context, picker *domain.CardPicker) error {
	names, err := p.ExpansionNameForCard(ctx, picker.CardName)
	if err != nil {
		return err
	}

	if picker.ExpansionName != "" && lo.Contains(names, picker.ExpansionName) {
		return nil
	}

	if len(names) == 0 {
		return ErrResultNotFound
	}

	if len(names) == 1 {
		picker.ExpansionName = names[0]
		return nil
	}

	question := &survey.Select{
		Message: "Select an expansion name:",
		Default: picker.ExpansionName,
		Options: names,
	}

	if err := survey.AskOne(question, &picker.ExpansionName); err != nil {
		return err
	}

	if exists, err := p.ExpansionNameExists(ctx, picker.ExpansionName); err != nil {
		return err
	} else if !exists {
		return ErrResultNotFound
	}

	return nil
}

var _ ports.ExpansionNameProvider = (*Prompt)(nil)
