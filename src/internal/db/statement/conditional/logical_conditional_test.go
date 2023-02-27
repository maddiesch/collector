package conditional_test

import (
	"database/sql"
	"testing"

	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/maddiesch/collector/internal/db/statement/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionalAnd(t *testing.T) {
	t.Run("simple equality", func(t *testing.T) {
		provider := generator.NewIncrementingArgumentNameProvider()

		stmt, args, err := conditional.And(
			conditional.Equal("First", 1),
			conditional.Equal("Second", 2),
		).Generate(provider)

		require.NoError(t, err)

		assert.Equal(t, `("First" = $v1 AND "Second" = $v2)`, stmt)

		if assert.Len(t, args, 2) {
			assert.Equal(t, sql.Named("v1", 1), args[0])
			assert.Equal(t, sql.Named("v2", 2), args[1])
		}
	})

	t.Run("nested equality", func(t *testing.T) {
		provider := generator.NewIncrementingArgumentNameProvider()

		stmt, args, err := conditional.And(
			conditional.Equal("First", 1),
			conditional.And(
				conditional.Equal("Second", 2),
				conditional.Equal("Third", 3),
			),
		).Generate(provider)

		require.NoError(t, err)

		assert.Equal(t, `("First" = $v1 AND ("Second" = $v2 AND "Third" = $v3))`, stmt)

		if assert.Len(t, args, 3) {
			assert.Equal(t, sql.Named("v1", 1), args[0])
			assert.Equal(t, sql.Named("v2", 2), args[1])
			assert.Equal(t, sql.Named("v3", 3), args[2])
		}
	})
}

func TestConditionalOr(t *testing.T) {
	provider := generator.NewIncrementingArgumentNameProvider()

	stmt, args, err := conditional.Or(
		conditional.Equal("First", 1),
		conditional.Equal("Second", 2),
	).Generate(provider)

	require.NoError(t, err)

	assert.Equal(t, `("First" = $v1 OR "Second" = $v2)`, stmt)

	if assert.Len(t, args, 2) {
		assert.Equal(t, sql.Named("v1", 1), args[0])
		assert.Equal(t, sql.Named("v2", 2), args[1])
	}
}
