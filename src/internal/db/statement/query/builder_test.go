package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Run("appends a ;", func(t *testing.T) {
		var b Builder

		b.WriteString("Foo")

		assert.Equal(t, "Foo;", b.String())
	})
}
