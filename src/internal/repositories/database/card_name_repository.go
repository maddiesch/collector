package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
)

func (d *Database) CardNameSearchPrefix(ctx context.Context, prefix string, expansion string) ([]string, error) {
	where := conditional.StringHasPrefix("Name", prefix)

	if expansion != "" {
		where = conditional.And(where, conditional.Equal("SetName", expansion))
	}

	stmt := statement.Select("Name").Distinct().From("Cache_DefaultCard").Where(where)

	rows, err := d.conn.QueryStatement(ctx, stmt)
	if err != nil {
		return nil, err
	}

	var names []string

	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			return nil, err
		}
		names = append(names, n)
	}

	return names, nil
}

var _ ports.CardNameRepository = (*Database)(nil)
