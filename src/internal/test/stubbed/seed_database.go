package stubbed

import (
	"encoding/json"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/stretchr/testify/require"
)

func LoadStubbedCardData(t require.TestingT) (cards []domain.CardData) {
	data, err := DataFS.ReadFile("data/bulk_data/default_cards.json")
	require.NoError(t, err)

	err = json.Unmarshal(data, &cards)
	require.NoError(t, err)

	return
}
