package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/maddiesch/collector/internal/test/stubbed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCardDataRepository(t *testing.T) {
	cards := stubbed.LoadStubbedCardData(t)

	t.Run("InsertCardData", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()

		err := db.InsertCardData(context.Background(), cards[0])

		assert.NoError(t, err)
	})

	t.Run("GetCardData", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()

		for _, card := range cards {
			err := db.InsertCardData(context.Background(), card)

			require.NoError(t, err)
		}

		t.Run("given a card that exists in the database", func(t *testing.T) {
			found, err := db.GetCardData(context.Background(), "Cloister Gargoyle", "Adventures in the Forgotten Realms", "302")

			require.NoError(t, err)

			if assert.Len(t, found, 1) {
				card := found[0]

				assert.Equal(t, "ab380bf5-2202-4098-b850-2cb660ec5351", card.ID)
				assert.Equal(t, "Cloister Gargoyle", card.Name)
				assert.Equal(t, "Adventures in the Forgotten Realms", card.SetName)
				assert.Equal(t, "en", card.LanguageTag)
				assert.Equal(t, "302", card.CollectorNumber)
				assert.Equal(t, 23, card.ReleasedAt.Day)
				assert.Equal(t, time.July, card.ReleasedAt.Month)
				assert.Equal(t, 2021, card.ReleasedAt.Year)
				assert.Equal(t, "{2}{W}", card.ManaCost)
				assert.Equal(t, true, card.HasFoil)
				assert.Equal(t, true, card.HasNormal)
				assert.Equal(t, "https://cards.scryfall.io/small/front/a/b/ab380bf5-2202-4098-b850-2cb660ec5351.jpg?1627711183", card.Image.Small)
				assert.Equal(t, "https://cards.scryfall.io/normal/front/a/b/ab380bf5-2202-4098-b850-2cb660ec5351.jpg?1627711183", card.Image.Normal)
				assert.Equal(t, "https://gatherer.wizards.com/Pages/Card/Details.aspx?multiverseid=530936", card.Links.Gatherer)
				assert.Equal(t, 10, card.Prices.NormalUSD)
				assert.Equal(t, 10, card.Prices.FoilUSD)
			}
		})
	})
}
