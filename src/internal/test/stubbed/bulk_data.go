package stubbed

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

func CreateBulkDataService(t require.TestingT) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/default-cards.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := DataFS.ReadFile("data/bulk_data/default_cards.json")
		require.NoError(t, err)

		io.Copy(w, bytes.NewReader(data))
	})

	return httptest.NewServer(mux)
}
