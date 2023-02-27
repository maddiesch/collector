package validate

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/samber/lo"

	"github.com/go-playground/validator/v10"
)

type NewInput struct {
	ports.ExpansionNameRepository
	ports.CardNameRepository
}

func New(in NewInput) ports.ValidationService {
	val := validator.New()

	val.RegisterValidationCtx("default_card_expansion", func(ctx context.Context, fl validator.FieldLevel) (r bool) {
		r, _ = in.ExpansionNameExists(ctx, fl.Field().String())
		return
	})

	val.RegisterValidationCtx("default_card_name", func(ctx context.Context, fl validator.FieldLevel) (r bool) {
		expansion, _, _, ok := fl.GetStructFieldOK2()
		if !ok {
			panic("Must provide the expansion parameter for the the card name e.g. `default_card_name=Expansion`")
		}

		r, _ = in.CardNameExists(ctx, fl.Field().String(), expansion.String())
		return
	})

	val.RegisterValidationCtx("card_condition", func(ctx context.Context, fl validator.FieldLevel) bool {
		return lo.Contains(domain.CardConditionAllString, fl.Field().String())
	})

	return &validate{
		Validate: val,
	}
}

type validate struct {
	*validator.Validate
}

var _ ports.ValidationService = (*validate)(nil)
