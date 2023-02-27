package database_test

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("Cleanup", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()

		err := db.Cleanup(context.Background())

		assert.NoError(t, err)
	})

	t.Run("Close", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		err := db.Close()

		assert.NoError(t, err)
	})
}
