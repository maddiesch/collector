package cache

import (
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/spf13/cobra"
)

func New(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cache",
		Short: "Manage the cache",
	}

	cmd.AddCommand(newUpdateCommand(config))

	return cmd
}
