package collection

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/color"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/command/config"
	"github.com/maddiesch/collector/cmd/collector-ctl/internal/cli/prompt"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/service/validate"
	"github.com/oklog/ulid/v2"
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
			validate := validate.New(validate.NewInput{
				ExpansionNameRepository: config.DB,
				CardNameRepository:      config.DB,
			})

			defaultExpansionName, err := p.ProvideDefaultExpansionName(cmd.Context())
			if errors.Is(err, prompt.ErrResultNotFound) {
				color.Warn.Println("Unable to find a default expansion name. Please try again.")
				return nil
			} else if err != nil {
				color.Error.Printf("Default expansion name provider error: %s\n", err)
				return err
			}

			groupName := ulid.Make().String()

		CardEntryLoop:
			for {
				picker := &domain.CardPicker{
					GroupName:     groupName,
					ExpansionName: defaultExpansionName,
				}

			RetryCardEntry:

				err := p.ProvideCardName(cmd.Context(), picker)
				if errors.Is(err, prompt.ErrResultNotFound) {
					if p.TryAgain("Unable to find card") {
						goto RetryCardEntry
					}
					color.Error.Println("Unable to find a card with that name. Please try again.")
					return nil
				} else if err != nil {
					color.Error.Printf("Card name provider error: %s\n", err)
					return err
				}

				err = p.ProvideExpansionName(cmd.Context(), picker)
				if errors.Is(err, prompt.ErrResultNotFound) {
					if p.TryAgain("Unable to find expansion name") {
						goto RetryCardEntry
					}
					color.Error.Println("Unable to find card in the specified expansion. Please try again.")
					return nil
				} else if err != nil {
					color.Error.Printf("Failed to get expansion name with error: %s\n", err)
					return err
				}

				err = p.ProvideCollectorNumber(cmd.Context(), picker)
				if errors.Is(err, prompt.ErrResultNotFound) {
					if p.TryAgain("Unable to find a collector number") {
						goto RetryCardEntry
					}
					color.Error.Println("Failed to select a collector number. Please try again.")
					return nil
				} else if err != nil {
					color.Error.Printf("Failed to select a collector number with error: %s\n", err)
					return err
				}

				err = p.ProvideCardFoil(cmd.Context(), picker)
				if err != nil {
					color.Error.Printf("Failed to select foil status: %s\n", err)
					return err
				}

				err = p.ProvideCardCondition(cmd.Context(), picker)
				if err != nil {
					color.Error.Printf("Failed to select a condition with error: %s\n", err)
					return err
				}

				err = p.ProvideLanguage(cmd.Context(), picker)
				if err != nil {
					color.Error.Printf("Failed to select a language with error: %s\n", err)
					return err
				}

				if err := validate.Struct(picker); err != nil {
					return err
				}

				card, err := p.ProvideCardDataFromPicker(cmd.Context(), picker)
				if err != nil {
					return err
				}

				save, finish := addCardCompleteAction(p)
				if save {
					collect := domain.CollectedCard{
						GroupName:       picker.GroupName,
						Name:            card.Name,
						SetName:         card.SetName,
						CollectorNumber: card.CollectorNumber,
						IsFoil:          picker.IsFoil,
						Condition:       picker.CardCondition,
						Language:        picker.Language,
						CreatedAt:       time.Now(),
					}

					if err := config.DB.InsertCollectedCard(cmd.Context(), collect); err != nil {
						return err
					}
				}
				if finish {
					break CardEntryLoop
				}
			}

			return nil
		},
	}
}

const (
	addCardSave     = "Save"
	addCardDiscard  = "Discard"
	addCardFinish   = "Finish"
	addCardContinue = "Continue"
)

func addCardCompleteAction(p *prompt.Prompt) (bool, bool) {
	saveCont := fmt.Sprintf("%s & %s", addCardSave, addCardContinue)
	saveFin := fmt.Sprintf("%s & %s", addCardSave, addCardFinish)
	discardCont := fmt.Sprintf("%s & %s", addCardDiscard, addCardContinue)
	discardFin := fmt.Sprintf("%s & %s", addCardDiscard, addCardFinish)

	question := &survey.Select{
		Message: "Next",
		Options: []string{saveCont, saveFin, discardCont, discardFin},
	}

	var selected string
	if err := survey.AskOne(question, &selected); err != nil {
		color.Error.Printf("Failed to select the next action: %s\n", err)
		return false, false
	}

	return strings.HasPrefix(selected, addCardSave), strings.HasSuffix(selected, addCardFinish)
}
