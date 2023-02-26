package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CardNameProvider interface {
	ProvideCardName(context.Context, *domain.CardPicker) error
}
