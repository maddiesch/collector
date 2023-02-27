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
		"Language":        data.LanguageTag,
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
		mapped, err := row.ReadMap()
		if err != nil {
			return err
		}

		var releasedAt domain.ReleaseDate

		if err := releasedAt.Scan(mapped["ReleasedAt"]); err != nil {
			return err
		}

		results = append(results, domain.CardData{
			ID:              mapped["ScryfallID"].(string),
			Name:            mapped["Name"].(string),
			SetName:         mapped["SetName"].(string),
			LanguageTag:     mapped["Language"].(string),
			CollectorNumber: mapped["CollectorNumber"].(string),
			ReleasedAt:      releasedAt,
			ManaCost:        mapped["ManaCost"].(string),
			HasFoil:         mapped["HasFoil"].(int64) == 1,
			HasNormal:       mapped["HasNormal"].(int64) == 1,
			Image: domain.CardDataImage{
				Small:  mapped["ImageSmallURL"].(string),
				Normal: mapped["ImageNormalURL"].(string),
			},
			Links: domain.CardDataLink{
				Gatherer: mapped["GathererURL"].(string),
			},
			Prices: domain.CardDataPrice{
				NormalUSD: int(mapped["PriceNormalUSD"].(int64)),
				FoilUSD:   int(mapped["PriceFoilUSD"].(int64)),
			},
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

var _ ports.CardDataRepository = (*Database)(nil)
