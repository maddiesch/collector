package cmd

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/maddiesch/collector/internal/magic"
	"github.com/spf13/cobra"
)

type CommandRuntime struct {
	WorkingDir string
}

func (r *CommandRuntime) CardDatabaseLocation() string {
	return filepath.Join(r.WorkingDir, "cache", "all_cards.sqlite")
}

func CreateRuntime(ctx context.Context, cmd *cobra.Command) (*CommandRuntime, error) {
	run := new(CommandRuntime)
	if path, err := cmd.Flags().GetString("path"); err == nil {
		if strings.HasPrefix(path, "~") {
			home, err := os.UserHomeDir()
			if err != nil {
				return nil, err
			}
			path = filepath.Join(home, path[1:])
		} else if !filepath.IsAbs(path) {
			path, err = filepath.Abs(path)
			if err != nil {
				return nil, err
			}
		}
		run.WorkingDir = path
	} else {
		return nil, err
	}

	return run, nil
}

func (r *CommandRuntime) NewCacheDB(ctx context.Context) (*magic.CardDB, error) {
	// TODO: Handle the case where there is no cache
	return magic.CreateCardDatabaseConn(ctx, magic.CreateCardDatabaseConnInput{
		FilePath: r.CardDatabaseLocation(),
	})
}
