package collection

import (
	"context"
	"encoding/csv"
	"strings"
	"testing"

	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/maddiesch/collector/internal/test/seed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_exportCollectionDeckbox(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()
		seed.Collection(t, db)

		var output strings.Builder

		err := exportCollectionDeckbox(context.Background(), exportCollectionInput{
			Database:  db,
			Output:    &output,
			Group:     "",
			ExportAll: true,
		})

		require.NoError(t, err)

		exported, err := csv.NewReader(strings.NewReader(output.String())).ReadAll()
		require.NoError(t, err)

		if assert.Len(t, exported, 6) {
			assert.Equal(t, "Count", exported[0][0])
			assert.Equal(t, "Name", exported[0][1])
			assert.Equal(t, "Edition", exported[0][2])
			assert.Equal(t, "Card Number", exported[0][3])
			assert.Equal(t, "Condition", exported[0][4])
			assert.Equal(t, "Language", exported[0][5])
			assert.Equal(t, "Foil", exported[0][6])
			assert.Equal(t, "Last Updated", exported[0][7])
		}
	})

	t.Run("last group", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()
		seed.Collection(t, db)

		var output strings.Builder

		err := exportCollectionDeckbox(context.Background(), exportCollectionInput{
			Database:  db,
			Output:    &output,
			Group:     "",
			ExportAll: false,
		})

		require.NoError(t, err)

		exported, err := csv.NewReader(strings.NewReader(output.String())).ReadAll()
		require.NoError(t, err)

		if assert.Len(t, exported, 2) {
			assert.Equal(t, "1", exported[1][0])
			assert.Equal(t, "Mountain", exported[1][1])
			assert.Equal(t, "Phyrexia: All Will Be One", exported[1][2])
			assert.Equal(t, "275", exported[1][3])
			assert.Equal(t, "Mint", exported[1][4])
			assert.Equal(t, "English", exported[1][5])
			assert.Equal(t, "", exported[1][6])
		}
	})

	t.Run("specific group", func(t *testing.T) {
		db := database.New(test.CreateDatabaseConnection(t))
		defer db.Close()
		seed.Collection(t, db)

		var output strings.Builder

		err := exportCollectionDeckbox(context.Background(), exportCollectionInput{
			Database:  db,
			Output:    &output,
			Group:     "named-import",
			ExportAll: false,
		})

		require.NoError(t, err)

		exported, err := csv.NewReader(strings.NewReader(output.String())).ReadAll()
		require.NoError(t, err)

		if assert.Len(t, exported, 2) {
			assert.Equal(t, "1", exported[1][0])
			assert.Equal(t, "Glissa Sunslayer", exported[1][1])
			assert.Equal(t, "Phyrexia: All Will Be One", exported[1][2])
			assert.Equal(t, "318", exported[1][3])
			assert.Equal(t, "Near Mint", exported[1][4])
			assert.Equal(t, "English", exported[1][5])
			assert.Equal(t, "", exported[1][6])
			assert.Equal(t, "2023-27-02 05:14:53 UTC", exported[1][7])
		}
	})
}
