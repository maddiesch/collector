package database

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/samber/lo"
	"golang.org/x/text/language"
)

func (d *Database) AvailableCachedCardLanguages(ctx context.Context) ([]string, error) {
	stmt := statement.Select("Language").Distinct().From("Cache_DefaultCard").OrderBy("Language", true)

	var results []string
	err := d.conn.EachRow(ctx, stmt, func(row *db.ResultRow) error {
		return db.ScanResultAppend(row.Rows, &results)
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(results, func(l string, _ int) string {
		tag, err := language.Parse(l)
		if err != nil {
			return l
		}

		spew.Dump(tag)

		return l
	}), nil
}

var _ ports.CardLanguageRepository = (*Database)(nil)
