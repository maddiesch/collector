package conditional

import (
	"database/sql"
	"fmt"

	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/maddiesch/collector/internal/db/statement/generator"
)

func StringLike(column, value string) Conditional {
	return &stringLikeConditional{column, value}
}

type stringLikeConditional struct {
	column string
	value  string
}

func (c *stringLikeConditional) Generate(p generator.ArgumentNameProvider) (string, []any, error) {
	name := p.Next()

	return fmt.Sprintf("%s LIKE $%s", dialect.Identifier(c.column), name), []any{sql.Named(name, c.value)}, nil
}

func StringHasPrefix(column, value string) Conditional {
	return StringLike(column, value+`%`)
}

func StringHasSuffix(column, value string) Conditional {
	return StringLike(column, `%`+value)
}
