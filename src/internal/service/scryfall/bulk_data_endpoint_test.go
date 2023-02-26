package scryfall_test

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/service/scryfall"
	"github.com/maddiesch/collector/internal/test/stubbed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBulkDataEndpoint(t *testing.T) {
	api := stubbed.CreateScryfallAPI(t)
	defer api.Close()

	service := scryfall.BulkDataEndpoint{
		BaseURL: api.URL,
	}

	t.Run(".GetBulkDataEndpoint", func(t *testing.T) {
		endpoints, err := service.GetBulkDataEndpoint(context.Background())

		require.NoError(t, err)

		assert.Len(t, endpoints, 5)
	})
}
