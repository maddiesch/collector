package magic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/maddiesch/collector/internal/raptor/raptortest"
	"github.com/maddiesch/collector/internal/task"
	"github.com/maddiesch/collector/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	cardDatabasePath string
)

func TestMain(t *testing.M) {
	cardDatabasePath = mock.CardDatabaseLocation()
	code := t.Run()
	os.Exit(code)
}

func TestUpdateCardDatabase(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v5/AllPrintings.sqlite.bz2" {
			http.Error(w, "not found", http.StatusNotFound)
		}

		filePath := mock.CompressedCardDatabasePath()

		file, err := os.Open(filePath)
		if err != nil {
			require.NoError(t, err)
		}
		defer file.Close()

		_, err = io.Copy(w, file)
		require.NoError(t, err)
	}))
	defer s.Close()

	t.Run("when downloading the compressed database file", func(t *testing.T) {
		downloadFile := filepath.Join(os.TempDir(), fmt.Sprintf("test-db-%d.sqlite", time.Now().UnixNano()))
		defer os.Remove(downloadFile)

		err := UpdateCardDatabase(context.Background(), UpdateCardDatabaseInput{
			DownloadTask:   new(task.NullTask),
			DecompressTask: new(task.NullTask),
			SourceURL:      s.URL + "/api/v5/AllPrintings.sqlite.bz2",
			FilePath:       downloadFile,
		})

		require.NoError(t, err)

		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: downloadFile,
		})

		require.NoError(t, err)
		defer conn.Close()

		_, err = conn.LastUpdatedAt(context.Background())
		assert.NoError(t, err)
	})

	t.Run("when updating to a directory that does not exist", func(t *testing.T) {
		downloadFile := filepath.Join(os.TempDir(), fmt.Sprintf("%d/test-db.sqlite", time.Now().UnixNano()))
		defer os.Remove(downloadFile)

		err := UpdateCardDatabase(context.Background(), UpdateCardDatabaseInput{
			DownloadTask:   new(task.NullTask),
			DecompressTask: new(task.NullTask),
			SourceURL:      s.URL + "/api/v5/AllPrintings.sqlite.bz2",
			FilePath:       downloadFile,
		})

		require.NoError(t, err)

		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: downloadFile,
		})

		require.NoError(t, err)
		defer conn.Close()

		_, err = conn.LastUpdatedAt(context.Background())
		assert.NoError(t, err)
	})
}

func TestCreateCardDatabaseConn(t *testing.T) {
	t.Run("given a database file that doesn't exist", func(t *testing.T) {
		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: "/foo/bar",
		})

		assert.ErrorIs(t, err, ErrDatabaseNotExists)
		assert.Nil(t, conn)
	})

	t.Run("given a db file that does exist", func(t *testing.T) {
		conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
			FilePath: cardDatabasePath,
		})
		require.NoError(t, err)
		defer conn.Close()
	})
}

func TestConn(t *testing.T) {
	conn, err := CreateCardDatabaseConn(context.Background(), CreateCardDatabaseConnInput{
		FilePath: cardDatabasePath,
	})
	require.NoError(t, err)
	defer conn.Close()

	conn.SetLogger(raptortest.NewQueryLogger(t))

	t.Run("LastUpdatedAt", func(t *testing.T) {
		lastUpdatedAt, err := conn.LastUpdatedAt(context.Background())
		require.NoError(t, err)

		assert.Equal(t, int64(1677632400), lastUpdatedAt.Unix())
	})
}
