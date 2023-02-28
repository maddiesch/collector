package collection

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func newExportCommand(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export the collection",
		RunE: func(cmd *cobra.Command, args []string) error {
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			group, err := cmd.Flags().GetString("group")
			if err != nil {
				return err
			}

			all, err := cmd.Flags().GetBool("all")
			if err != nil {
				return err
			}

			if fn, ok := exportCollectionFormats[format]; !ok {
				color.Error.Printf(`Unsupported format "%s" valid options: (%s)`+"\n", format, strings.Join(lo.Keys(exportCollectionFormats), ", "))
				return nil
			} else {
				return fn(cmd.Context(), exportCollectionInput{
					Database:  config.DB,
					Output:    cmd.OutOrStdout(),
					Group:     group,
					ExportAll: all,
				})
			}
		},
	}

	cmd.PersistentFlags().StringP("group", "g", "", "Select a collection group to export. Default to the last import group")
	cmd.PersistentFlags().Bool("all", false, "Export the entire collection")
	cmd.PersistentFlags().StringP("format", "f", "deckbox", "Select the format for the exported collection")

	return cmd
}

type exportCollectionInput struct {
	*database.Database

	Output io.Writer

	Group     string
	ExportAll bool
}

func (i exportCollectionInput) groupName(ctx context.Context) string {
	if i.ExportAll {
		return ""
	}
	if i.Group == "" {
		var group string
		found, _ := i.Database.GetMetadata(ctx, database.MetadataInsertCollectedLastAddedGroup, &group)
		if found {
			return group
		}
	}
	return i.Group
}

type exportCollectionFunc func(context.Context, exportCollectionInput) error

var exportCollectionFormats = map[string]exportCollectionFunc{
	"deckbox":  exportCollectionDeckbox,
	"decklist": exportCollectionDecklist,
	"echomtg":  exportCollectionEchoMTG,
}

func exportCollectionDecklist(ctx context.Context, in exportCollectionInput) error {
	cards, err := in.Database.GetCollectedCards(ctx, in.groupName(ctx))
	if err != nil {
		return err
	}

	list := map[string]int{}

	for _, card := range cards {
		if current, ok := list[card.Name]; ok {
			list[card.Name] = current + 1
		} else {
			list[card.Name] = 1
		}
	}

	for name, count := range list {
		fmt.Fprintf(in.Output, "%d %s\n", count, name)
	}

	return nil
}

func exportCollectionEchoMTG(ctx context.Context, in exportCollectionInput) error {
	cards, err := in.Database.GetCollectedCards(ctx, in.groupName(ctx))
	if err != nil {
		return err
	}

	w := csv.NewWriter(in.Output)
	if err := w.Write([]string{"Quantity", "Name", "Set", "CardNumber", "Condition", "Foil"}); err != nil {
		return err
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Name < cards[j].Name
	})

	for _, card := range cards {
		foil := "False"
		if card.IsFoil {
			foil = "True"
		}

		// TODO: Map condition values
		fields := []string{"1", card.Name, card.SetName, card.CollectorNumber, "NM", foil}

		if err := w.Write(fields); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func exportCollectionDeckbox(ctx context.Context, in exportCollectionInput) error {
	cards, err := in.Database.GetCollectedCards(ctx, in.groupName(ctx))
	if err != nil {
		return err
	}

	w := csv.NewWriter(in.Output)
	if err := w.Write([]string{"Count", "Name", "Edition", "Card Number", "Condition", "Language", "Foil", "Last Updated"}); err != nil {
		return err
	}

	for _, card := range cards {
		var foil string
		if card.IsFoil {
			foil = "foil"
		}
		lang := card.Language
		if lang == "Phyrexian" {
			lang = "English"
		}

		updatedAt := card.UpdatedAt.UTC().Format("2006-02-01 15:04:05 MST")
		fields := []string{"1", card.Name, card.SetName, card.CollectorNumber, string(card.Condition), lang, foil, updatedAt}

		if err := w.Write(fields); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}
