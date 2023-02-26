package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
)

func (d *Database) CollectorNumberForCard(ctx context.Context, cardName, expansionName string) ([]string, error) {
	stmt := statement.Select("CollectorNumber").Distinct().From("Cache_DefaultCard").Where(
		conditional.And(
			conditional.Equal("Name", cardName),
			conditional.Equal("SetName", expansionName),
		),
	).OrderBy("CollectorNumber", true)

	var results []string

	err := d.conn.EachRow(ctx, stmt, func(r *db.ResultRow) error {
		return db.ScanResultAppend(r.Rows, &results)
	})

	return results, err
}

var _ ports.CollectorNumberRepository = (*Database)(nil)
