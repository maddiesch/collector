package db_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/maddiesch/collector/internal/db"
	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConn(t *testing.T) {
	conn, err := db.NewConn(context.Background(), db.NewConnInput{
		FilePath:     t.Name(),
		IsMemoryMode: true,
	})

	require.NoError(t, err)
	assert.NoError(t, conn.Close())
}

func TestConnTransaction(t *testing.T) {
	conn, err := db.NewConn(context.Background(), db.NewConnInput{
		FilePath:     t.Name(),
		IsMemoryMode: true,
	})
	require.NoError(t, err)
	defer conn.Close()

	t.Run("commits if there is no error", func(t *testing.T) {
		err := conn.Transaction(context.Background(), func(tx *db.Tx) error {
			return tx.ExecStatement(context.Background(), statement.Insert().Into("Metadata").ValueMap(map[string]any{
				"Key":   t.Name(),
				"Value": "committed",
			}))
		})
		require.NoError(t, err)

		var v string
		err = conn.QueryStatementRow(context.Background(), statement.Select("Value").From("Metadata").Where(conditional.Equal("Key", t.Name())).Limit(1)).Scan(&v)
		require.NoError(t, err)

		assert.Equal(t, "committed", v)
	})

	var (
		errTriggerRollback = errors.New("trigger rollback")
	)

	t.Run("rolls back if there is an error", func(t *testing.T) {
		err := conn.Transaction(context.Background(), func(tx *db.Tx) error {
			err := tx.ExecStatement(context.Background(), statement.Insert().Into("Metadata").ValueMap(map[string]any{
				"Key":   t.Name(),
				"Value": "this should not exist",
			}))
			require.NoError(t, err)

			return errTriggerRollback
		})
		require.ErrorIs(t, err, errTriggerRollback)

		var v string
		err = conn.QueryStatementRow(context.Background(), statement.Select("Value").From("Metadata").Where(conditional.Equal("Key", t.Name())).Limit(1)).Scan(&v)
		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.Equal(t, "", v)
	})

	t.Run("rolls back if the func panics", func(t *testing.T) {
		assert.Panics(t, func() {
			conn.Transaction(context.Background(), func(tx *db.Tx) error {
				err := tx.ExecStatement(context.Background(), statement.Insert().Into("Metadata").ValueMap(map[string]any{
					"Key":   t.Name(),
					"Value": "this should not exist",
				}))

				require.NoError(t, err)

				panic("trigger a rollback")
			})
		})

		var v string
		err = conn.QueryStatementRow(context.Background(), statement.Select("Value").From("Metadata").Where(conditional.Equal("Key", t.Name())).Limit(1)).Scan(&v)
		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.Equal(t, "", v)
	})
}
