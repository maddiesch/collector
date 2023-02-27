package prompt

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/service/lang"
	"github.com/samber/lo"
)

func (p *Prompt) ProvideCardDataFromPicker(ctx context.Context, picker *domain.CardPicker) (*domain.CardData, error) {
	result, err := p.GetCardData(ctx, picker.CardName, picker.ExpansionName, picker.CollectorNumber)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, ErrResultNotFound
	}
	if len(result) == 1 {
		return lo.ToPtr(result[0]), nil
	}

	question := &survey.Select{
		Message: "Found multiple results, please select one:",
		Options: lo.Map(result, func(r domain.CardData, i int) string {
			return fmt.Sprintf("(%d) %s - %s #%s [%s]", i+1, r.Name, r.SetName, r.CollectorNumber, lang.DisplayNameForTagString(r.LanguageTag))
		}),
	}

	var selected string
	if err := survey.AskOne(question, &selected); err != nil {
		return nil, err
	}

	indexRegex := regexp.MustCompile(`\A\((\d+)\)`)

	results := indexRegex.FindAllStringSubmatch(selected, -1)
	if len(results) != 1 {
		return nil, errors.New("unable to parse index from the selection (t), this should't technically happen")
	}
	if len(results[0]) != 2 {
		return nil, errors.New("unable to parse index from the selection (i), this should't technically happen")
	}

	index, err := strconv.ParseInt(results[0][1], 10, 32)
	if err != nil {
		return nil, errors.New("unable to parse index from the selection (p), this should't technically happen")
	}

	return lo.ToPtr(result[index-1]), nil
}

var _ ports.CardDataProvider = (*Prompt)(nil)
