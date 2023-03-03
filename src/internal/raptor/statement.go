package raptor

import (
	"context"

	"github.com/maddiesch/collector/internal/raptor/statement/generator"
)

func (c *Conn) QueryRowStatement(ctx context.Context, statement generator.Generator) *Row {
	query, args, err := statement.Generate()
	if err != nil {
		return &Row{err: err}
	}
	return c.QueryRow(ctx, query, args...)
}
