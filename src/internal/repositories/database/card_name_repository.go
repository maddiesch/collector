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

	stmt := statement.Select("Name").Distinct().From("Cache_DefaultCard").Where(where).OrderBy("SetName", true)

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

func (d *Database) CardNameExists(ctx context.Context, name string, expansion string) (bool, error) {
	where := conditional.Equal("Name", name)
	if expansion != "" {
		where = conditional.And(where, conditional.Equal("SetName", expansion))
	}

	stmt := statement.Exists(
		statement.Select("1").From("Cache_DefaultCard").Where(where),
	)

	var exists bool
	if err := d.conn.QueryStatementRow(ctx, stmt).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

var _ ports.CardNameRepository = (*Database)(nil)
