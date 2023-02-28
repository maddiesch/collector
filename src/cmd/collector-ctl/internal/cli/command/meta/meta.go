package meta

import (
	"fmt"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/output"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/spf13/cobra"
)

func New(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meta",
		Short: "Manage metadata about the collection",
	}

	cmd.AddCommand(newLastInsertedGroup(config))

	return cmd
}

func newLastInsertedGroup(config config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "last-inserted-group",
		Short: "Get the name of the last inserted group",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := output.New(cmd.OutOrStdout(), cmd.ErrOrStderr())

			var name string

			found, err := config.DB.GetMetadata(cmd.Context(), database.MetadataInsertCollectedLastAddedGroup, &name)
			if err != nil {
				out.Error("Metadata lookup failed with error: %s", err)
				return err
			}

			if found {
				fmt.Fprintln(cmd.OutOrStdout(), name)
			} else {
				out.Warn("There doesn't seem to be a last inserted group")
			}

			return nil
		},
	}
}
