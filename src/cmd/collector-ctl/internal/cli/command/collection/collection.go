package collection

import (
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/spf13/cobra"
)

func New(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "collection",
		Short:   "Mange collection of cards",
		Aliases: []string{"col", "c"},
	}

	cmd.AddCommand(newAddCommand(config))
	cmd.AddCommand(newExportCommand(config))
	cmd.AddCommand(newAddBulkCommand(config))

	return cmd
}
