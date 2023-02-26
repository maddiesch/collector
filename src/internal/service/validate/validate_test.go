package validate

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCustomValidation(t *testing.T) {
	repo := &repository{
		cards: map[string][]string{
			"Phyrexia: All Will Be One": {"Mountain", "Plains"},
		},
	}

	validate := New(NewInput{
		CardNameRepository:      repo,
		ExpansionNameRepository: repo,
	})

	t.Run("default_card_expansion", func(t *testing.T) {
		t.Run("given a valid expansion name", func(t *testing.T) {
			object := struct {
				Name string `validate:"required,default_card_expansion"`
			}{"Phyrexia: All Will Be One"}

			err := validate.Struct(object)

			assert.NoError(t, err)
		})
	})

	t.Run("default_card_name", func(t *testing.T) {
		t.Run("given a valid expansion name and card name", func(t *testing.T) {
			object := struct {
				Expansion string `validate:"required,default_card_expansion"`
				Name      string `validate:"required,default_card_name=Expansion"`
			}{"Phyrexia: All Will Be One", "Mountain"}

			err := validate.Struct(object)

			assert.NoError(t, err)
		})
	})
}

type repository struct {
	ports.CardNameRepository
	ports.ExpansionNameRepository

	cards map[string][]string
}

func (r *repository) ExpansionNameExists(_ context.Context, e string) (bool, error) {
	return lo.Contains(lo.Keys(r.cards), e), nil
}

func (r *repository) CardNameExists(_ context.Context, n, e string) (bool, error) {
	if cards, ok := r.cards[e]; ok {
		return lo.Contains(cards, n), nil
	} else {
		return false, nil
	}
}
