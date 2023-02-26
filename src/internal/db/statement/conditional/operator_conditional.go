package conditional

import (
	"database/sql"
	"fmt"

	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/maddiesch/collector/internal/db/statement/generator"
)

func Equal(column string, value any) Conditional {
	return &operatorConditional{column, "=", value}
}

type operatorConditional struct {
	column   string
	operator string
	value    any
}

func (c *operatorConditional) Generate(provider generator.ArgumentNameProvider) (string, []any, error) {
	name := provider.Next()

	return fmt.Sprintf("%s %s $%s", dialect.Identifier(c.column), c.operator, name), []any{sql.Named(name, c.value)}, nil
}
