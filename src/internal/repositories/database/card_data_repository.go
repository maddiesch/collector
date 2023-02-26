package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db/statement"
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

var _ ports.CardDataRepository = (*Database)(nil)
