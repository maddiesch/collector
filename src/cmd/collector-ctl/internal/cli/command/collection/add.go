package collection

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/prompt"
	"github.com/spf13/cobra"
)

func newAddCommand(config config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "add-interactive",
		Short:   "Add a new card to the collection",
		Aliases: []string{"add"},
		RunE: func(cmd *cobra.Command, args []string) error {
			p := &prompt.Prompt{
				Database: config.DB,
			}

			defaultExpansionName, err := p.ProvideExpansionName(cmd.Context(), "")
			if errors.Is(err, prompt.ErrResultNotFound) {
				color.Warn.Println("Unable to find a default expansion name. Please try again.")
				return nil
			} else if err != nil {
				color.Error.Printf("Default expansion name provider error: %s\n", err)
				return err
			}

			spew.Dump(defaultExpansionName)

			return nil
		},
	}
}
