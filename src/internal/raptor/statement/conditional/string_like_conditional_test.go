package conditional_test

import (
	"database/sql"
	"testing"

	"github.com/maddiesch/collector/internal/raptor/statement/conditional"
	"github.com/maddiesch/collector/internal/raptor/statement/generator"
	"github.com/stretchr/testify/assert"
)

func TestStringHasPrefix(t *testing.T) {
	provider := generator.NewIncrementingArgumentNameProvider()

	str, args := conditional.StringHasPrefix("Foo", "Bar").Generate(provider)

	assert.Equal(t, `"Foo" LIKE $v1`, str)

	if assert.Len(t, args, 1) {
		assert.Equal(t, sql.Named("v1", "Bar%"), args[0])
	}
}

func TestStringHasSuffix(t *testing.T) {
	provider := generator.NewIncrementingArgumentNameProvider()

	str, args := conditional.StringHasSuffix("Foo", "Bar").Generate(provider)

	assert.Equal(t, `"Foo" LIKE $v1`, str)

	if assert.Len(t, args, 1) {
		assert.Equal(t, sql.Named("v1", "%Bar"), args[0])
	}
}
