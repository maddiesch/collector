package prompt

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
)

func (p *Prompt) ProvideCollectorNumber(ctx context.Context, picker *domain.CardPicker) error {
	numbers, err := p.CollectorNumberForCard(ctx, picker.CardName, picker.ExpansionName)
	if err != nil {
		return err
	}
	if len(numbers) == 0 {
		return ErrResultNotFound
	}
	if len(numbers) == 1 {
		picker.CollectorNumber = numbers[0]
		return nil
	}

	question := &survey.Select{
		Message: "Collector Number:",
		Help:    "Found in the lower left-hand corner of the card. ",
		Options: numbers,
	}

	if err := survey.AskOne(question, &picker.CollectorNumber); err != nil {
		return err
	}

	return nil
}

var _ ports.CollectorNumberProvider = (*Prompt)(nil)
