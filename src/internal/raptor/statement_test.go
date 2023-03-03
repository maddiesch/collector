package raptor_test

import (
	"testing"

	"github.com/maddiesch/collector/internal/raptor/internal/test"
	"github.com/maddiesch/collector/internal/raptor/statement"
	"github.com/maddiesch/collector/internal/raptor/statement/conditional"
	"github.com/stretchr/testify/assert"
)

func TestConn_QueryRowStatement(t *testing.T) {
	conn, ctx := test.Setup(t)
	defer conn.Close()

	query := statement.Select("FirstName", "LastName").From("People").Where(conditional.Equal("FirstName", "Maddie")).Limit(1)

	var firstName, lastName string
	err := conn.QueryRowStatement(ctx, query).Scan(&firstName, &lastName)

	assert.NoError(t, err)

	assert.Equal(t, "Maddie", firstName)
	assert.Equal(t, "Schipper", lastName)
}
