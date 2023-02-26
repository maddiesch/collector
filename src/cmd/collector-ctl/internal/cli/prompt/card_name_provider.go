package prompt

import (
	"context"

	"github.com/maddiesch/collector/internal/core/ports"
)

func (p *Prompt) ProvideCardName(ctx context.Context, expansion string) (string, error) {
	return "", ErrResultNotFound
}

var _ ports.CardNameProvider = (*Prompt)(nil)
