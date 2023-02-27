package dialect_test

import (
	"testing"

	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	assert.Equal(t, `"Foo"`, dialect.Identifier("Foo"))
}

func TestIdentifierMap(t *testing.T) {
	id := lo.Map([]string{"Foo", "Bar"}, dialect.IdentifierMap)

	assert.Equal(t, `"Foo"`, id[0])
	assert.Equal(t, `"Bar"`, id[1])
}
