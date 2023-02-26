package statement_test

import (
	"database/sql"
	"testing"

	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectBuilder(t *testing.T) {
	t.Run("with a where clause", func(t *testing.T) {
		query, args, err := statement.Select().From("TestTable").Where(conditional.Equal("Key", "foo")).Generate()

		require.NoError(t, err)

		assert.Equal(t, `SELECT * FROM "TestTable" WHERE "Key" = $v1;`, query)

		if assert.Len(t, args, 1) {
			assert.Equal(t, sql.Named("v1", "foo"), args[0])
		}
	})

	t.Run("with selected columns", func(t *testing.T) {
		query, _, err := statement.Select("FirstName", "LastName").From("TestTable").Generate()

		require.NoError(t, err)

		assert.Equal(t, `SELECT "FirstName", "LastName" FROM "TestTable";`, query)
	})

	t.Run("with a distinct limit", func(t *testing.T) {
		query, _, err := statement.Select().Distinct().From("TestTable").Limit(1).Generate()

		require.NoError(t, err)

		assert.Equal(t, `SELECT DISTINCT * FROM "TestTable" LIMIT 1;`, query)
	})
}
