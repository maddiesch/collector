package mock_test

import (
	"os"
	"path/filepath"
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

func TestCopyCardDatabase(t *testing.T) {
	path := filepath.Join(os.TempDir(), "card_database.sqlite")
	defer os.Remove(path)

	err := mock.CopyCardDatabase(path)

	assert.NoError(t, err)
}
