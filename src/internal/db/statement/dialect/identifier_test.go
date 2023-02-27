package dialect_test

import (
	"testing"

	"github.com/maddiesch/collector/internal/db/statement/dialect"
	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	assert.Equal(t, `"Foo"`, dialect.Identifier("Foo"))
}
