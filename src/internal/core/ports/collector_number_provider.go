package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CollectorNumberProvider interface {
	ProvideCollectorNumber(context.Context, *domain.CardPicker) error
}
