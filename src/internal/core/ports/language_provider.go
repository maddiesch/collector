package ports

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
)

type LanguageProvider interface {
	ProvideLanguage(context.Context, *domain.CardPicker) error
}
