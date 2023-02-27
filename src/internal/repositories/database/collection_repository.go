package database

import (
	"context"
	"time"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
)

const (
	MetadataInsertCollectedLastAddedGroup = "collection.last_added_group"
)

func (d *Database) InsertCollectedCard(ctx context.Context, card domain.CollectedCard) error {
	stmt := statement.Insert().Into("Inventory").ValueMap(map[string]any{
		"ID":              card.ID,
		"GroupName":       card.GroupName,
		"Name":            card.Name,
		"SetName":         card.SetName,
		"CollectorNumber": card.CollectorNumber,
		"Condition":       card.Condition,
		"Language":        card.Language,
		"IsFoil":          card.IsFoil,
		"CreatedAt":       card.CreatedAt.Unix(),
		"UpdatedAt":       card.CreatedAt.Unix(),
	})

	return d.conn.Transaction(ctx, func(tx *db.Tx) error {
		if err := tx.ExecStatement(ctx, stmt); err != nil {
			return err
		}

		return setMetadataTx(ctx, tx, MetadataInsertCollectedLastAddedGroup, card.GroupName)
	})
}

func (d *Database) GetCollectedCards(ctx context.Context, groupName string) ([]domain.CollectedCard, error) {
	var where conditional.Conditional
	if groupName != "" {
		where = conditional.Equal("GroupName", groupName)
	}

	stmt := statement.Select().From("Inventory").Where(where).OrderBy("UpdatedAt", false)

	var cards []domain.CollectedCard

	err := d.conn.EachRow(ctx, stmt, func(r *db.ResultRow) error {
		result, err := r.ReadMap()
		if err != nil {
			return err
		}

		cards = append(cards, domain.CollectedCard{
			ID:              result["ID"].(string),
			GroupName:       result["GroupName"].(string),
			Name:            result["Name"].(string),
			SetName:         result["SetName"].(string),
			CollectorNumber: result["CollectorNumber"].(string),
			IsFoil:          result["IsFoil"].(int64) == 1,
			Language:        result["Language"].(string),
			Condition:       domain.CardCondition(result["Condition"].(string)),
			CreatedAt:       time.Unix(result["CreatedAt"].(int64), 0),
			UpdatedAt:       time.Unix(result["CreatedAt"].(int64), 0),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return cards, nil
}

var _ ports.CollectionRepository = (*Database)(nil)
