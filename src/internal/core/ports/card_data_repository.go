package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CardDataRepository interface {
	FlushCardData(context.Context) error

	InsertCardData(context.Context, domain.CardData) error
}
