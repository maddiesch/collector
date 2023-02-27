package database

import (
	"context"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db/statement"
)

func (d *Database) InsertCollectedCard(ctx context.Context, card domain.CollectedCard) error {
	stmt := statement.Insert().Into("Inventory").ValueMap(map[string]any{
		"GroupName":       card.GroupName,
		"Name":            card.Name,
		"SetName":         card.SetName,
		"CollectorNumber": card.CollectorNumber,
		"Condition":       card.Condition,
		"Language":        card.Language,
		"IsFoil":          card.IsFoil,
		"CreatedAt":       card.CreatedAt.Unix(),
	})

	if err := d.conn.ExecStatement(ctx, stmt); err != nil {
		return err
	}

	return d.SetMetadata(ctx, "collection.last_added_group", card.GroupName)
}

var _ ports.CollectionRepository = (*Database)(nil)
