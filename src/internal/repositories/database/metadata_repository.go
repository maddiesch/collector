package database

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/pkg/errors"
)

func (d *Database) SetMetadata(ctx context.Context, key string, value any) error {
	return d.conn.Transaction(ctx, func(tx *db.Tx) error {
		return setMetadataTx(ctx, tx, key, value)
	})
}

func setMetadataTx(ctx context.Context, tx *db.Tx, key string, value any) error {
	vData, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal value into json")
	}

	stmt := statement.Insert().OrReplace().Into("Metadata").ValueMap(map[string]any{
		"Key":   key,
		"Value": vData,
	})

	return tx.ExecStatement(ctx, stmt)
}

func (d *Database) GetMetadata(ctx context.Context, key string, target any) (bool, error) {
	stmt := statement.Select("Value").From("Metadata").Where(conditional.Equal("Key", key))

	var value []byte
	if err := d.conn.QueryStatementRow(ctx, stmt).Scan(&value); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, json.Unmarshal(value, target)
}

func (d *Database) DeleteMetadata(ctx context.Context, key string) error {
	return d.conn.Transaction(ctx, func(tx *db.Tx) error {
		return deleteMetadataTx(ctx, tx, key)
	})
}

func deleteMetadataTx(ctx context.Context, tx *db.Tx, key string) error {
	stmt := statement.Delete().From("Metadata").Where(conditional.Equal("Key", key))
	return tx.ExecStatement(ctx, stmt)
}

var _ ports.MetadataRepository = (*Database)(nil)
