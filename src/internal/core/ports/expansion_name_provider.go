package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type ExpansionNameProvider interface {
	ProvideExpansionName(context.Context, *domain.CardPicker) error
}
