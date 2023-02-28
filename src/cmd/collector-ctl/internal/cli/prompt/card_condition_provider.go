package prompt

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/samber/lo"
)

func (p *Prompt) ProvideCardCondition(_ context.Context, picker *domain.CardPicker) error {
	if lo.Contains(domain.CardConditionAll, picker.CardCondition) {
		return nil
	}
	question := &survey.Select{
		Message: "Card Condition",
		Default: string(domain.CardConditionMint),
		Options: domain.CardConditionAllString,
	}

	var strCondition string

	if err := survey.AskOne(question, &strCondition); err != nil {
		return err
	}

	if lo.Contains(domain.CardConditionAllString, strCondition) {
		picker.CardCondition = domain.CardCondition(strCondition)
		return nil
	} else {
		return ErrResultNotFound
	}
}

var _ ports.CardConditionProvider = (*Prompt)(nil)
