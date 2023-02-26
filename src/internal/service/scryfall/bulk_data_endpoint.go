package scryfall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
)

type BulkDataEndpoint struct {
	BaseURL string
}

func (s *BulkDataEndpoint) GetBulkDataEndpoint(ctx context.Context) ([]domain.BulkDataEndpoint, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", s.BaseURL+"/bulk-data", nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code: %d", response.StatusCode)
	}

	defer response.Body.Close()

	body := struct {
		Data []domain.BulkDataEndpoint `json:"data"`
	}{}

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body.Data, nil
}

var _ ports.BulkDataEndpointRepository = (*BulkDataEndpoint)(nil)
