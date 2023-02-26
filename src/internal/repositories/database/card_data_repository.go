package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
)

func (d *Database) InsertCardData(ctx context.Context, data domain.CardData) error {
	stmt := statement.Insert().Into("Cache_DefaultCard").ValueMap(map[string]any{
		"ScryfallID":      data.ID,
		"Name":            data.Name,
		"SetName":         data.SetName,
		"CollectorNumber": data.CollectorNumber,
		"Language":        data.Language,
		"ReleasedAt":      data.ReleasedAt,
		"ImageSmallURL":   data.Image.Small,
		"ImageNormalURL":  data.Image.Normal,
		"ManaCost":        data.ManaCost,
		"HasFoil":         data.HasFoil,
		"HasNormal":       data.HasNormal,
		"GathererURL":     data.Links.Gatherer,
		"PriceNormalUSD":  data.Prices.NormalUSD,
		"PriceFoilUSD":    data.Prices.FoilUSD,
	})

	return d.conn.ExecStatement(ctx, stmt)
}

func (d *Database) FlushCardData(ctx context.Context) error {
	return d.conn.ExecStatement(ctx, statement.Delete().From("Cache_DefaultCard"))
}

func (d *Database) GetCardData(ctx context.Context, cardName, expansionName, collectorNumber string) ([]domain.CardData, error) {
	where := conditional.And(
		conditional.Equal("Name", cardName),
		conditional.And(
			conditional.Equal("SetName", expansionName),
			conditional.Equal("CollectorNumber", collectorNumber),
		),
	)

	stmt := statement.Select().From("Cache_DefaultCard").Where(where)

	var results []domain.CardData

	err := d.conn.EachRow(ctx, stmt, func(row *db.ResultRow) error {
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

var _ ports.CardDataRepository = (*Database)(nil)
