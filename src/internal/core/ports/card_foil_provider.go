package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CardFoilProvider interface {
	ProvideCardFoil(context.Context, *domain.CardPicker) error
}
