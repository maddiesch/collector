package magic

import (
	"context"
	"os"
	"testing"

	"github.com/maddiesch/collector/internal/raptor/raptortest"
	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	cardDatabasePath string
)

func TestMain(t *testing.M) {
	cardDatabasePath = mock.CardDatabaseLocation()
	code := t.Run()
	os.Exit(code)
}

func TestCreateCardDatabaseConn(t *testing.T) {
	t.Run("given a database file that doesn't exist", func(t *testing.T) {
		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: "/foo/bar",
		})

		assert.ErrorIs(t, err, ErrDatabaseNotExists)
		assert.Nil(t, conn)
	})

	t.Run("given a db file that does exist", func(t *testing.T) {
		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: cardDatabasePath,
		})
		require.NoError(t, err)
		defer conn.Close()
	})
}

func TestConn(t *testing.T) {
	conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
		FilePath: cardDatabasePath,
	})
	require.NoError(t, err)
	defer conn.Close()

	conn.SetLogger(raptortest.NewQueryLogger(t))

	t.Run("LastUpdatedAt", func(t *testing.T) {
		lastUpdatedAt, err := conn.LastUpdatedAt(context.Background())
		require.NoError(t, err)

		assert.Equal(t, int64(1677632400), lastUpdatedAt.Unix())
	})
}
