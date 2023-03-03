package mock

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateCardDatabaseServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v5/AllPrintings.sqlite.bz2", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		filePath := CompressedCardDatabasePath()
		file, err := os.Open(filePath)
		require.NoError(t, err)
		defer file.Close()

		_, err = io.Copy(w, file)
		require.NoError(t, err)
	}))

	return httptest.NewServer(mux)
}
