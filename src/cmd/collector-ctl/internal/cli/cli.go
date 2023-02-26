package cli

import (
	"context"
	"os"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context, config config.Config) {
	root := command.NewRootCommand(config)

	status := execute(ctx, root, config)

	os.Exit(status)
}

func execute(ctx context.Context, cmd *cobra.Command, config config.Config) int {
	defer func() {
		if err := config.Shutdown(ctx); err != nil {
			color.Warn.Printf("Shutdown error: %s\n", err)
		}
	}()

	_, err := cmd.ExecuteContextC(ctx)

	if err != nil {
		color.Error.Print("Error: ")
		color.Red.Println(err.Error())
		return 1
	}

	return 0
}
