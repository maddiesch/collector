package command

import (
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/cache"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/collection"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/spf13/cobra"
)

func NewRootCommand(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "commander-ctl",
	}

	cmd.AddCommand(newVersionCommand())
	cmd.AddCommand(cache.New(config))
	cmd.AddCommand(collection.New(config))

	return cmd
}
