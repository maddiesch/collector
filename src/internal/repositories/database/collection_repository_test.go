package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestCollectionRepository(t *testing.T) {
	t.Run("InsertCollectedCard", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()

		err := db.InsertCollectedCard(context.Background(), domain.CollectedCard{
			GroupName:       "TestGroup",
			Name:            "Mountain",
			SetName:         "Phyrexia: All Will Be One",
			CollectorNumber: "270",
			IsFoil:          true,
			Language:        "English",
			Condition:       domain.CardConditionMint,
			CreatedAt:       time.Now(),
		})

		assert.NoError(t, err)
	})
}