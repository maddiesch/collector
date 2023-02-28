package collection

import (
	"time"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/prompt"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/output"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/service/validate"
	"github.com/oklog/ulid/v2"
	"github.com/spf13/cobra"
)

func newAddBulkCommand(config config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-bulk",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := output.New(cmd.OutOrStdout(), cmd.ErrOrStderr())

			picker := new(domain.CardPicker)

			if value, err := cmd.Flags().GetString("name"); err != nil {
				return err
			} else if value == "" {
				out.Error("Must provide a name")
				return nil
			} else {
				picker.CardName = value
			}
			if value, err := cmd.Flags().GetString("group"); err != nil {
				return err
			} else if value == "" {
				picker.GroupName = ulid.Make().String()
			} else {
				picker.GroupName = value
			}
			if value, err := cmd.Flags().GetString("expansion"); err != nil {
				return err
			} else if value == "" {
				out.Error("Must provide an expansion")
				return nil
			} else {
				picker.ExpansionName = value
			}
			if value, err := cmd.Flags().GetString("collector-number"); err != nil {
				return err
			} else if value == "" {
				out.Error("Must provide a collector-number")
				return nil
			} else {
				picker.CollectorNumber = value
			}
			if value, err := cmd.Flags().GetString("quality"); err != nil {
				return err
			} else if value == "" {
				out.Error("Must provide a quality")
				return nil
			} else {
				picker.CardCondition = domain.CardCondition(value)
			}
			if value, err := cmd.Flags().GetString("language"); err != nil {
				return err
			} else if value == "" {
				out.Error("Must provide a language")
				return nil
			} else {
				picker.Language = value
			}
			if value, err := cmd.Flags().GetBool("foil"); err != nil {
				return err
			} else {
				picker.IsFoil = value
			}

			count, err := cmd.Flags().GetInt("count")
			if err != nil {
				return err
			}

			validate := validate.New(validate.NewInput{
				ExpansionNameRepository: config.DB,
				CardNameRepository:      config.DB,
			})

			if err := validate.Struct(picker); err != nil {
				return err
			}

			p := &prompt.Prompt{
				Database: config.DB,
			}

			card, err := p.ProvideCardDataFromPicker(cmd.Context(), picker)
			if err != nil {
				return err
			}

			for i := 0; i < count; i++ {
				collect := domain.CollectedCard{
					ID:              ulid.Make().String(),
					ScryfallID:      card.ID,
					GroupName:       picker.GroupName,
					Name:            card.Name,
					SetName:         card.SetName,
					CollectorNumber: card.CollectorNumber,
					IsFoil:          picker.IsFoil,
					Condition:       picker.CardCondition,
					Language:        picker.Language,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				}

				if err := config.DB.InsertCollectedCard(cmd.Context(), collect); err != nil {
					return err
				}
			}

			return nil
		},
	}

	// TODO: Make the descriptions better
	cmd.PersistentFlags().StringP("name", "n", "", "the name of the cards to bulk add to the collection")
	cmd.PersistentFlags().StringP("group", "g", "", "The group the bulk import cards will be added to")
	cmd.PersistentFlags().StringP("expansion", "e", "", "the expansion of the bulk import cards")
	cmd.PersistentFlags().StringP("collector-number", "c", "", "the collector number of the card to import")
	cmd.PersistentFlags().StringP("quality", "q", "", "the quality of the card to be imported")
	cmd.PersistentFlags().StringP("language", "l", "", "the language of the card to be imported")
	cmd.PersistentFlags().Bool("foil", false, "the cards to import are foil")
	cmd.PersistentFlags().Int("count", 1, "The number of cards to be added to the collection")

	return cmd
}
