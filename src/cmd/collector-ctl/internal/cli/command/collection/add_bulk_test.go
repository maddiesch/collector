package collection

import (
	"context"
	"strings"
	"testing"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/maddiesch/collector/internal/test/seed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newAddBulkCommand(t *testing.T) {
	db := database.New(test.CreateDatabaseConnection(t))
	defer db.Close()

	seed.DefaultCardCache(t, db)

	t.Run("given valid arguments", func(t *testing.T) {
		var stdout, stderr strings.Builder

		cmd := newAddBulkCommand(config.Config{
			DB: db,
		})
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)
		cmd.SetArgs([]string{
			"--group", "Test Bulk Import",
			"--name", "Phyrexian Colossus",
			"--expansion", "The List",
			"--collector-number", "949",
			"--quality", "Near Mint",
			"--language", "English",
			"--count", "16",
		})

		err := cmd.Execute()

		require.NoError(t, err)

		cards, err := db.GetCollectedCards(context.Background(), "")

		require.NoError(t, err)

		assert.Len(t, cards, 16)
	})
}
