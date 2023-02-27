package domain

import "github.com/samber/lo"

type CardCondition string

const (
	CardConditionMint          CardCondition = "Mint"
	CardConditionNearMint      CardCondition = "Near Mint"
	CardConditionGood          CardCondition = "Good (Lightly Played)"
	CardConditionPlayed        CardCondition = "Played"
	CardConditionHeavilyPlayed CardCondition = "Heavily Played"
	CardConditionPoor          CardCondition = "Poor"
)

var (
	CardConditionAll = []CardCondition{
		CardConditionMint,
		CardConditionNearMint,
		CardConditionGood,
		CardConditionPlayed,
		CardConditionHeavilyPlayed,
		CardConditionPoor,
	}

	CardConditionAllString = lo.Map(CardConditionAll, func(c CardCondition, _ int) string {
		return string(c)
	})
)
