package mock_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCardDatabaseServer(t *testing.T) {
	server := mock.CreateCardDatabaseServer(t)
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/api/v5/AllPrintings.sqlite.bz2", nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
