package cache

import (
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/report"
	"github.com/maddiesch/collector/internal/service/cache"
	"github.com/maddiesch/collector/internal/service/cdatstream"
	"github.com/maddiesch/collector/internal/service/scryfall"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func newUpdateCommand(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the cached card database",
		RunE: func(cmd *cobra.Command, args []string) error {
			reporter := &report.ProgressBarReporter{
				Bar: progressbar.Default(-1),
			}

			service := cache.UpdateCardCacheService{
				BulkDataEndpointRepository: &scryfall.BulkDataEndpoint{
					BaseURL: "https://api.scryfall.com",
				},
				CardDataStreamer:   cdatstream.New(),
				CardDataRepository: config.DB,
				MetadataRepository: config.DB,
				ProgressReporter:   reporter,
			}

			reporter.Draw()

			return service.Call(cmd.Context())
		},
	}

	return cmd
}
