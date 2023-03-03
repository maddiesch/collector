package mock_test

import (
	"strings"
	"testing"

	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestCompressedCardDatabasePath(t *testing.T) {
	path := mock.CompressedCardDatabasePath()

	assert.True(t, strings.HasSuffix(path, "/src/internal/test/mock/card_database.sqlite.bz2"))
}

func TestCardDatabaseLocation(t *testing.T) {
	assert.NotPanics(t, func() {
		mock.CardDatabaseLocation()
	})
}
