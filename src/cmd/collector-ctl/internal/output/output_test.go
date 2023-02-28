package output_test

import (
	"strings"
	"testing"

	"github.com/maddiesch/collector/cmd/collector-ctl/internal/output"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	var stdout, stderr strings.Builder

	out := output.New(&stdout, &stderr)

	t.Run("Print", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.Print(t.Name())

			assert.True(t, strings.HasSuffix(stdout.String(), "\n"), "the stdout does not have a trailing newline")
			assert.Equal(t, "", stderr.String())
		})
	})

	t.Run("Println", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.Println(t.Name())

			assert.True(t, strings.HasSuffix(stdout.String(), "\n"), "the stdout does not have a trailing newline")
			assert.Equal(t, "", stderr.String())
		})
	})

	t.Run("PrintE", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.PrintE(t.Name())

			assert.True(t, strings.HasSuffix(stderr.String(), "\n"), "the stderr does not have a trailing newline")
			assert.Equal(t, "", stdout.String())
		})
	})

	t.Run("PrintEln", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.PrintEln(t.Name())

			assert.True(t, strings.HasSuffix(stderr.String(), "\n"), "the stderr does not have a trailing newline")
			assert.Equal(t, "", stdout.String())
		})
	})

	t.Run("Warn", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.Warn(t.Name())

			assert.True(t, strings.HasSuffix(stderr.String(), "\n"), "the stderr does not have a trailing newline")
			assert.Equal(t, "", stdout.String())
		})
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.Error(t.Name())

			assert.True(t, strings.HasSuffix(stderr.String(), "\n"), "the stderr does not have a trailing newline")
			assert.Equal(t, "", stdout.String())
		})
	})

	t.Run("Success", func(t *testing.T) {
		t.Run("appends a newline", func(t *testing.T) {
			defer stdout.Reset()
			defer stderr.Reset()

			out.Success(t.Name())

			assert.True(t, strings.HasSuffix(stderr.String(), "\n"), "the stderr does not have a trailing newline")
			assert.Equal(t, "", stdout.String())
		})
	})
}
