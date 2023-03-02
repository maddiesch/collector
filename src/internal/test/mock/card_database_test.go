package mock_test

import (
	"testing"

	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestCardDatabaseLocation(t *testing.T) {
	assert.NotPanics(t, func() {
		mock.CardDatabaseLocation()
	})
}
