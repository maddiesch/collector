package seed

import (
	"context"
	"encoding/json"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test/stubbed"
	"github.com/stretchr/testify/require"
)

func DefaultCardCache(t require.TestingT, db *database.Database) {
	data, err := stubbed.DataFS.ReadFile("data/bulk_data/default_cards.json")

	require.NoError(t, err, "Failed to load the default cards")

	var cards []domain.CardData

	err = json.Unmarshal(data, &cards)
	require.NoError(t, err)

	for _, card := range cards {
		err = db.InsertCardData(context.Background(), card)
		require.NoError(t, err)
	}
}
