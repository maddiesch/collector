package test

import (
	"context"

	"github.com/maddiesch/collector/internal/db"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func CreateDatabaseConnection(t require.TestingT) *db.Conn {
	conn, err := db.NewConn(context.Background(), db.NewConnInput{
		FilePath:     "testing-db-" + ulid.Make().String(),
		IsMemoryMode: true,
	})
	require.NoError(t, err)

	return conn
}
