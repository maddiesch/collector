package prompt

import (
	"context"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/service/lang"
	"github.com/samber/lo"
)

func (p *Prompt) ProvideLanguage(ctx context.Context, picker *domain.CardPicker) error {
	cards, err := p.GetCardData(ctx, picker.CardName, picker.ExpansionName, picker.CollectorNumber)
	if err != nil {
		return err
	}
	if len(cards) == 0 {
		return ErrResultNotFound
	}

	names := lo.Map(cards, func(card domain.CardData, _ int) string {
		return lang.DisplayNameForTagString(card.LanguageTag)
	})
	sort.Strings(names)

	switch len(names) {
	case 0:
		input := &survey.Input{
			Message: "Card Language",
		}
		return survey.AskOne(input, &picker.Language)
	case 1:
		picker.Language = names[0]
		return nil
	default:
		question := &survey.Select{
			Message: "Select Language:",
			Options: names,
		}
		return survey.AskOne(question, &picker.Language)
	}
}

var _ ports.LanguageProvider = (*Prompt)(nil)
