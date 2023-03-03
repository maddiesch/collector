package cmd

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/maddiesch/collector/internal/test"
	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRuntime(t *testing.T) {
	server := mock.CreateCardDatabaseServer(t)
	defer server.Close()

	conf := config{
		SourceURL: server.URL + "/api/v5/AllPrintings.sqlite.bz2",
	}

	t.Run("when the database cache does not exist", func(t *testing.T) {
		workingDir := test.CreateWorkingDir(t)
		defer os.RemoveAll(workingDir)

		var stdout, stderr strings.Builder

		cmd := newRootCommand(conf)
		cmd.SetArgs([]string{"cache", "age", "--path", workingDir})
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)

		err := cmd.Execute()

		assert.Error(t, err)

		assert.Contains(t, stderr.String(), "card database does not exist")
	})

	t.Run("when the database cache does exist", func(t *testing.T) {
		workingDir := test.CreateWorkingDir(t)
		defer os.RemoveAll(workingDir)

		t.Run("call update command", func(t *testing.T) {
			cmd := newRootCommand(conf)
			cmd.SetArgs([]string{"cache", "update", "--path", workingDir})
			cmd.SetOut(io.Discard)
			cmd.SetErr(io.Discard)

			err := cmd.Execute()

			require.NoError(t, err)
		})

		t.Run("returns the last updated time", func(t *testing.T) {
			var stdout, stderr strings.Builder

			cmd := newRootCommand(conf)
			cmd.SetArgs([]string{"cache", "age", "--path", workingDir})
			cmd.SetOut(&stdout)
			cmd.SetErr(&stderr)

			err := cmd.Execute()

			require.NoError(t, err)

			assert.Equal(t, "2023-02-28\n", stdout.String())
		})
	})
}
