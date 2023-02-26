package database

import (
	"context"
	"encoding/json"

	"github.com/maddiesch/collector/internal/core/ports"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/pkg/errors"
)

func (d *Database) SetMetadata(ctx context.Context, key string, value any) error {
	vData, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal value into json")
	}

	stmt := statement.Insert().OrReplace().Into("Metadata").ValueMap(map[string]any{
		"Key":   key,
		"Value": vData,
	})

	return d.conn.ExecStatement(ctx, stmt)
}

var _ ports.MetadataRepository = (*Database)(nil)
