package prompt

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
)

func (p *Prompt) ProvideCardName(ctx context.Context, picker *domain.CardPicker) error {
	input := &survey.Input{
		Message: "Card Name:",
		Help:    "The name is found at the top of the physical card.",
		Default: picker.CardName,
		Suggest: func(toComplete string) []string {
			names, _ := p.CardNameSearchPrefix(ctx, toComplete, picker.ExpansionName)
			return names
		},
	}

Retry:
	if err := survey.AskOne(input, &picker.CardName); err != nil {
		return err
	}

	if picker.CardName == "" {
		color.Warn.Println("Please input a card name\n")
		goto Retry
	}

	if found, err := p.CardNameExists(ctx, picker.CardName, picker.ExpansionName); err != nil {
		return err
	} else if !found && picker.ExpansionName == "" {
		if p.TryAgain("Unable to find a card with that name") {
			goto Retry
		}
		return ErrResultNotFound
	} else if !found {
		message := fmt.Sprintf("The card (%s) does not exist in %s, try again without the expansion name?", picker.CardName, picker.ExpansionName)
		if p.Confirm(message, true) {
			picker.ExpansionName = ""
			goto Retry
		}
		return ErrResultNotFound
	}

	return nil
}

var _ ports.CardNameProvider = (*Prompt)(nil)
