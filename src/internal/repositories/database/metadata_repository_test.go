package database_test

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataRepository(t *testing.T) {
	db := database.New(test.CreateDatabaseConnection(t))
	defer db.Close()

	t.Run("SetMetadata", func(t *testing.T) {
		err := db.SetMetadata(context.Background(), "testing.set-value", "my test value")

		assert.NoError(t, err)
	})

	t.Run("GetMetadata", func(t *testing.T) {
		type object struct{ Name string }

		val := object{"MTG"}

		err := db.SetMetadata(context.Background(), "testing.object-value", val)

		require.NoError(t, err)

		t.Run("loading a value that exists", func(t *testing.T) {
			var load object

			found, err := db.GetMetadata(context.Background(), "testing.object-value", &load)

			require.NoError(t, err)

			assert.True(t, found)
			assert.Equal(t, "MTG", load.Name)
		})

		t.Run("loading a value does not exist", func(t *testing.T) {
			var load object

			found, err := db.GetMetadata(context.Background(), "testing.object-missing", &load)

			require.NoError(t, err)

			assert.False(t, found)
			assert.Equal(t, "", load.Name)
		})
	})

	t.Run("DeleteMetadata", func(t *testing.T) {
		err := db.SetMetadata(context.Background(), "testing.object-delete", "test-deleted")
		require.NoError(t, err)

		err = db.DeleteMetadata(context.Background(), "testing.object-delete")
		require.NoError(t, err)

		found, err := db.GetMetadata(context.Background(), "testing.object-delete", nil)
		require.NoError(t, err)

		assert.False(t, found)
	})
}
