package data_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maddiesch/collector/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	_ "embed"
)

//go:embed example.json.bz2
var bzip2ExampleData []byte

func createDownloadServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(bzip2ExampleData); err != nil {
			panic(err)
		}
	}))
}

func TestDownload(t *testing.T) {
	server := createDownloadServer()
	defer server.Close()

	t.Run("download file", func(t *testing.T) {
		err := data.Download(context.Background(), data.DownloadInput{
			Logger:  zaptest.NewLogger(t),
			FromURL: server.URL,
			Dest:    io.Discard,
		})

		assert.NoError(t, err)
	})
}

func TestInflateCompressedFile(t *testing.T) {
	var output bytes.Buffer

	err := data.InflateCompressedFile(context.Background(), bytes.NewReader(bzip2ExampleData), &output)

	require.NoError(t, err)

	content := struct {
		TestFile string
	}{}

	err = json.Unmarshal(output.Bytes(), &content)
	require.NoError(t, err)

	assert.Equal(t, "Compressed as bz2", content.TestFile)
}
