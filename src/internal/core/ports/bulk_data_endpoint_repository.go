package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type BulkDataEndpointRepository interface {
	GetBulkDataEndpoint(context.Context) ([]domain.BulkDataEndpoint, error)
}
