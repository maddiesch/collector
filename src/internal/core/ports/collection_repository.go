package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type CollectionRepository interface {
	InsertCollectedCard(context.Context, domain.CollectedCard) error
}