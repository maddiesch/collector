package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/oklog/ulid/v2"
)

func CreateWorkingDir(t *testing.T) string {
	path := filepath.Join(mock.TempDir(), ulid.Make().String())
	if err := os.MkdirAll(path, 0766); err != nil {
		if t == nil {
			panic(err)
		}
		t.Fatalf("failed to create working dir: %v", err)
	}

	return path
}
