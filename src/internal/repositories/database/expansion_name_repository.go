package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
)

func (d *Database) ExpansionNameSearchPrefix(ctx context.Context, prefix string, name string) ([]string, error) {
	where := conditional.StringHasPrefix("SetName", prefix)

	if name != "" {
		where = conditional.And(where, conditional.Equal("Name", name))
	}

	stmt := statement.Select("SetName").Distinct().From("Cache_DefaultCard").Where(where)

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

func (d *Database) ExpansionNameExists(ctx context.Context, name string) (bool, error) {
	stmt := statement.Exists(
		statement.Select("1").From("Cache_DefaultCard").Where(conditional.Equal("SetName", name)),
	)

	var exists bool
	err := d.conn.QueryStatementRow(ctx, stmt).Scan(&exists)

	return exists, err
}

var _ ports.ExpansionNameRepository = (*Database)(nil)
