package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWorkingDir(t *testing.T) {
	path := CreateWorkingDir(t)
	defer os.RemoveAll(path)

	assert.DirExists(t, path)
}
