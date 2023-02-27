package statement

import (
	"database/sql"
	"strings"

	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/maddiesch/collector/internal/db/statement/generator"
	"github.com/maddiesch/collector/internal/db/statement/query"
)

type InsertValue struct {
	ColumnName string
	Value      any
}

type InsertBuilder struct {
	tableName string
	orReplace bool
	values    map[string]any
}

func Insert() *InsertBuilder {
	return &InsertBuilder{
		values: make(map[string]any),
	}
}

func (b *InsertBuilder) Into(tableName string) *InsertBuilder {
	b.tableName = tableName

	return b
}

func (b *InsertBuilder) OrReplace() *InsertBuilder {
	b.orReplace = true

	return b
}

func (b *InsertBuilder) ValueMap(m map[string]any) *InsertBuilder {
	for k, v := range m {
		b.values[k] = v
	}

	return b
}

func (b *InsertBuilder) Value(column string, value any) *InsertBuilder {
	b.values[column] = value

	return b
}

func (b *InsertBuilder) Generate() (string, []any, error) {
	var query query.Builder
	var args []any

	query.WriteString("INSERT ")
	if b.orReplace {
		query.WriteString("OR REPLACE ")
	}
	query.WriteStringf("INTO %s ", dialect.Identifier(b.tableName))

	provider := generator.NewIncrementingArgumentNameProvider()

	if len(b.values) == 0 {
		query.WriteString("DEFAULT VALUES")
	} else {
		var columns, values []string

		for column, value := range b.values {
			vName := provider.Next()
			columns = append(columns, dialect.Identifier(column))
			values = append(values, "$"+vName)
			args = append(args, sql.Named(vName, value))
		}

		query.WriteStringf("(%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(values, ", "))
	}

	return query.String(), args, nil
}
