package conditional_test

import (
	"database/sql"
	"testing"

	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/maddiesch/collector/internal/db/statement/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringHasPrefix(t *testing.T) {
	provider := generator.NewIncrementingArgumentNameProvider()

	str, args, err := conditional.StringHasPrefix("Foo", "Bar").Generate(provider)

	require.NoError(t, err)

	assert.Equal(t, `"Foo" LIKE $v1`, str)

	if assert.Len(t, args, 1) {
		assert.Equal(t, sql.Named("v1", "Bar%"), args[0])
	}
}

func TestStringHasSuffix(t *testing.T) {
	provider := generator.NewIncrementingArgumentNameProvider()

	str, args, err := conditional.StringHasSuffix("Foo", "Bar").Generate(provider)

	require.NoError(t, err)

	assert.Equal(t, `"Foo" LIKE $v1`, str)

	if assert.Len(t, args, 1) {
		assert.Equal(t, sql.Named("v1", "%Bar"), args[0])
	}
}
