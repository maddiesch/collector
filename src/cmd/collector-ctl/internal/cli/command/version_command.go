package command

import (
	"github.com/maddiesch/collector/internal/core"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Commander")
			cmd.Printf("  Core: %s\n", core.Version)
		},
	}
}
