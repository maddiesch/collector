package stubbed

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func CreateScryfallAPI(t require.TestingT) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/bulk-data", func(w http.ResponseWriter, r *http.Request) {
		data, err := DataFS.ReadFile("data/scryfall_api/bulk-data.json")
		require.NoError(t, err)

		io.Copy(w, bytes.NewReader(data))
	})

	return httptest.NewServer(mux)
}
