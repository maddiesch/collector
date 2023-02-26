package cdatstream

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/test/stubbed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamCardData(t *testing.T) {
	server := stubbed.CreateBulkDataService(t)
	defer server.Close()

	service := New()

	stream, err := service.StreamCardData(context.Background(), server.URL+"/default-cards.json")

	require.NoError(t, err)

	var count int

	for event := range stream {
		switch event := event.(type) {
		case error:
			require.NoError(t, event)
		case domain.CardData:
			count++
		default:
			require.FailNowf(t, "Invalid event type", "Type: %T", event)
		}
	}

	assert.Equal(t, 48, count)
}
