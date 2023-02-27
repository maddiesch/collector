package conditional

import (
	"database/sql"
	"fmt"

	"github.com/maddiesch/collector/internal/raptor/statement/dialect"
	"github.com/maddiesch/collector/internal/raptor/statement/generator"
)

func Equal(column string, value any) Conditional {
	return &operatorInfixConditional{column, "=", value}
}

func NotEqual(col string, val any) Conditional {
	return &operatorInfixConditional{col, "!=", val}
}

func LessThan(col string, val any) Conditional {
	return &operatorInfixConditional{col, "<", val}
}

func LessThanEq(col string, val any) Conditional {
	return &operatorInfixConditional{col, "<=", val}
}

func GreaterThan(col string, val any) Conditional {
	return &operatorInfixConditional{col, ">", val}
}

func GreaterThannEq(col string, val any) Conditional {
	return &operatorInfixConditional{col, ">=", val}
}

type operatorInfixConditional struct {
	column   string
	operator string
	value    any
}

func (c *operatorInfixConditional) Generate(provider generator.ArgumentNameProvider) (string, []any) {
	name := provider.Next()

	return fmt.Sprintf("%s %s $%s", dialect.Identifier(c.column), c.operator, name), []any{sql.Named(name, c.value)}
}
