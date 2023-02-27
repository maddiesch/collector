package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CardConditionProvider interface {
	ProvideCardCondition(context.Context, *domain.CardPicker) error
}
