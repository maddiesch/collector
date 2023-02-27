package prompt

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
)

func (p *Prompt) ProvideCardFoil(ctx context.Context, picker *domain.CardPicker) error {
	cards, err := p.GetCardData(ctx, picker.CardName, picker.ExpansionName, picker.CollectorNumber)
	if err != nil {
		return err
	}

	var hasFoil, hasNormal bool

	if len(cards) == 1 {
		hasFoil = cards[0].HasFoil
		hasNormal = cards[0].HasNormal
	} else {
		hasFoil = true
		hasNormal = true
	}

	if !hasNormal {
		picker.IsFoil = true
		return nil
	}

	if !hasFoil {
		picker.IsFoil = false
		return nil
	}

	question := &survey.Confirm{
		Message: "Foil card?",
		Default: false,
	}
	return survey.AskOne(question, &picker.IsFoil)
}

var _ ports.CardFoilProvider = (*Prompt)(nil)
