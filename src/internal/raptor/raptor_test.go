package raptor_test

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/raptor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

var (
	ConnString = "file:testing-db?mode=memory&cache=shared"
)

type TestingT interface {
	zaptest.TestingT
	require.TestingT
}

func createTestConnection(t TestingT) (*raptor.Conn, context.Context) {
	ctx := context.Background()
	conn, err := raptor.New(ConnString)
	require.NoError(t, err)

	conn.SetLogger(zaptest.NewLogger(t))

	_, err = conn.Exec(ctx, `CREATE TABLE "TestTable" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Name" TEXT NOT NULL DEFAULT '', "Age" INTEGER NOT NULL DEFAULT 0);`)
	require.NoError(t, err)

	return conn, ctx
}

func TestNewConn(t *testing.T) {
	conn, err := raptor.New(ConnString)
	require.NoError(t, err)

	assert.NoError(t, conn.Close())
}

func TestConn_Transact(t *testing.T) {
	conn, ctx := createTestConnection(t)
	defer conn.Close()

	t.Run("insert", func(t *testing.T) {
		err := conn.Transact(ctx, func(tx raptor.DB) error {
			_, err := tx.Exec(ctx, `INSERT INTO "TestTable" DEFAULT VALUES;`)
			return err
		})

		assert.NoError(t, err)
	})

	t.Run("nested", func(t *testing.T) {
		err := conn.Transact(ctx, func(tx raptor.DB) error {
			_, err := tx.Exec(ctx, `INSERT INTO "TestTable" DEFAULT VALUES;`)
			require.NoError(t, err)

			return tx.Transact(ctx, func(tx raptor.DB) error {
				_, err := tx.Exec(ctx, `INSERT INTO "TestTable" DEFAULT VALUES;`)
				return err
			})
		})

		assert.NoError(t, err)
	})
}
