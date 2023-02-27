package db_test

import (
	"context"
	"testing"

	"github.com/maddiesch/collector/internal/db"
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
