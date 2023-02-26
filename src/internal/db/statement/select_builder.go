package statement

import (
	"strings"

	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/maddiesch/collector/internal/db/statement/generator"
	"github.com/maddiesch/collector/internal/db/statement/query"
	"github.com/samber/lo"
)

type SelectBuilder struct {
	tableName  string
	isDistinct bool
	columns    []string
	where      conditional.Conditional
	limit      *int64
}

func Select(columns ...string) *SelectBuilder {
	return &SelectBuilder{
		columns: columns,
	}
}

func (b *SelectBuilder) From(table string) *SelectBuilder {
	b.tableName = table

	return b
}

func (b *SelectBuilder) Distinct() *SelectBuilder {
	b.isDistinct = true

	return b
}

func (b *SelectBuilder) Where(condition conditional.Conditional) *SelectBuilder {
	b.where = condition

	return b
}

func (b *SelectBuilder) Limit(l int64) *SelectBuilder {
	b.limit = lo.ToPtr(l)

	return b
}

func (b *SelectBuilder) Generate() (string, []any, error) {
	var query query.Builder
	var args []any

	query.WriteString("SELECT ")

	if b.isDistinct {
		query.WriteString("DISTINCT ")
	}

	if len(b.columns) == 0 {
		query.WriteRune('*')
	} else {
		query.WriteString(strings.Join(lo.Map(b.columns, dialect.IdentifierMap), ", "))
	}

	query.WriteStringf(" FROM %s", dialect.Identifier(b.tableName))

	provider := generator.NewIncrementingArgumentNameProvider()

	if b.where != nil {
		where, wArgs, err := b.where.Generate(provider)
		if err != nil {
			return "", nil, err
		}

		query.WriteString(" WHERE ")
		query.WriteString(where)

		args = append(args, wArgs...)
	}

	if b.limit != nil {
		query.WriteStringf(" LIMIT %d", lo.FromPtr(b.limit))
	}

	return query.String(), args, nil
}

var _ generator.Generator = (*SelectBuilder)(nil)
