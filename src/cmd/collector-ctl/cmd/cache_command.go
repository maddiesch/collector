package cmd

import (
	"fmt"

	"github.com/maddiesch/collector/internal/magic"
	"github.com/maddiesch/collector/internal/task"
	"github.com/spf13/cobra"
)

const (
	defaultCardCacheURL = "https://mtgjson.com/api/v5/AllPrintings.sqlite.bz2"
)

func newCacheCommand(c config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cache",
		Short: "Manage the local card cache",
	}

	cmd.AddCommand(newCacheUpdateCommand(c))
	cmd.AddCommand(newCacheAgeCommand())

	return cmd
}

func newCacheAgeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "age",
		Short: "Get the age of the local card cache",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			run, err := CreateRuntime(cmd.Context(), cmd)
			if err != nil {
				return err
			}

			db, err := run.NewCacheDB(cmd.Context())
			if err != nil {
				return err
			}

			updatedAt, err := db.LastUpdatedAt(cmd.Context())
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", updatedAt.Format("2006-01-02"))

			return err
		},
	}

	return cmd
}

func newCacheUpdateCommand(c config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the local card cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			run, err := CreateRuntime(cmd.Context(), cmd)
			if err != nil {
				return err
			}

			return magic.UpdateCardDatabase(cmd.Context(), magic.UpdateCardDatabaseInput{
				DownloadTask:   new(task.NullTask), // TODO: Report Progress
				DecompressTask: new(task.NullTask), // TODO: Report Progress
				SourceURL:      c.SourceURL,
				FilePath:       run.CardDatabaseLocation(),
			})
		},
	}

	return cmd
}
