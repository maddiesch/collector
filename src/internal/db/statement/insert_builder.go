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

/**

type InsertValue struct {
	Column string
	Value  any
}

type InsertBuilder interface {
	OrReplace() InsertBuilder
	Into(string) InsertBuilder
	ValueMap(map[string]any) InsertBuilder
	Value(string, any) InsertBuilder
	Values([]InsertValue) InsertBuilder

	Generate() (string, []any, error)
}

func Insert() InsertBuilder {
	return &insertBuilder{}
}

type insertBuilder struct {
	orReplace bool
	tableName string
	values    []InsertValue
}

func (i *insertBuilder) Into(tableName string) InsertBuilder {
	i.tableName = tableName
	return i
}

func (i *insertBuilder) OrReplace() InsertBuilder {
	i.orReplace = true
	return i
}

func (i *insertBuilder) ValueMap(v map[string]any) InsertBuilder {
	var values []InsertValue

	for k, v := range v {
		values = append(values, InsertValue{k, v})
	}

	return i.Values(values)
}

func (i *insertBuilder) Value(c string, v any) InsertBuilder {
	return i.Values([]InsertValue{{Column: c, Value: v}})
}

func (i *insertBuilder) Values(values []InsertValue) InsertBuilder {
	i.values = append(i.values, values...)

	return i
}

func (i *insertBuilder) Generate() (string, []any, error) {
	var builder strings.Builder

	builder.WriteString("INSERT ")
	if i.orReplace {
		builder.WriteString("OR REPLACE ")
	}
	builder.WriteString("INTO ")
	builder.WriteString(escapeString(i.tableName))

	var args []any
	var columns, values []string

	for i, v := range i.values {
		columns = append(columns, escapeString(v.Column))
		name := fmt.Sprintf("v%d", i+1)
		values = append(values, ":"+name)
		args = append(args, sql.Named(name, v.Value))
	}

	builder.WriteString(" (")
	builder.WriteString(strings.Join(columns, ", "))
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(");")

	return builder.String(), args, nil
}

*/
