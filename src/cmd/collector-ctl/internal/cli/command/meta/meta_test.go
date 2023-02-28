package meta

import (
	"context"
	"strings"
	"testing"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/maddiesch/collector/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newLastInsertedGroup(t *testing.T) {
	db := database.New(test.CreateDatabaseConnection(t))
	defer db.Close()

	setup := func(t *testing.T) (*cobra.Command, *strings.Builder, *strings.Builder) {
		stdout := new(strings.Builder)
		stderr := new(strings.Builder)

		cmd := newLastInsertedGroup(config.Config{
			DB: db,
		})
		cmd.SetOut(stdout)
		cmd.SetErr(stderr)

		return cmd, stdout, stderr
	}

	t.Run("when there is no last inserted group", func(t *testing.T) {
		cmd, stdout, stderr := setup(t)

		err := cmd.Execute()

		require.NoError(t, err)

		assert.Equal(t, "", stdout.String())
		assert.Contains(t, stderr.String(), "last inserted group")
	})

	t.Run("when there is a last inserted group", func(t *testing.T) {
		err := db.SetMetadata(context.Background(), database.MetadataInsertCollectedLastAddedGroup, t.Name())
		require.NoError(t, err)

		cmd, stdout, stderr := setup(t)

		err = cmd.Execute()
		require.NoError(t, err)

		assert.Equal(t, t.Name()+"\n", stdout.String())
		assert.Equal(t, "", stderr.String())
	})
}
