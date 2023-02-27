package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CardDataProvider interface {
	ProvideCardDataFromPicker(context.Context, *domain.CardPicker) (*domain.CardData, error)
}
