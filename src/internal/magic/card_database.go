package magic

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/maddiesch/collector/internal/data"
	"github.com/maddiesch/collector/internal/raptor"
	"github.com/maddiesch/collector/internal/raptor/statement"
	"github.com/maddiesch/collector/internal/task"
	"github.com/mattn/go-sqlite3"
)

var (
	ErrDatabaseNotExists = errors.New("card database does not exist")
)

type CardDB struct {
	*raptor.Conn
}

type CreateCardDatabaseConnInput struct {
	FilePath string
}

func CreateCardDatabaseConn(ctx context.Context, in CreateCardDatabaseConnInput) (*CardDB, error) {
	connStr := fmt.Sprintf("file:%s?mode=ro", in.FilePath)

	conn, err := raptor.New(connStr)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrCantOpen {
			return nil, ErrDatabaseNotExists
		}
		return nil, err
	}

	return &CardDB{
		Conn: conn,
	}, nil
}

type UpdateCardDatabaseInput struct {
	SourceURL string // https://mtgjson.com/api/v5/AllPrintings.sqlite.bz2
	FilePath  string

	DownloadTask   task.Task
	DecompressTask task.Task
}

func UpdateCardDatabase(ctx context.Context, in UpdateCardDatabaseInput) error {
	tempDownload, err := os.CreateTemp("", "card_db_update-*.sqlite.bz2")
	if err != nil {
		return err
	}
	defer os.Remove(tempDownload.Name())
	defer tempDownload.Close()

	err = data.Download(ctx, data.DownloadInput{
		Task:    in.DownloadTask,
		FromURL: in.SourceURL,
		Dest:    tempDownload,
	})
	if err != nil {
		return err
	}

	if _, err := tempDownload.Seek(0, io.SeekStart); err != nil {
		return err
	}

	dbFile, err := os.CreateTemp("", "card_db_update-*.sqlite")
	if err != nil {
		return err
	}
	defer os.Remove(dbFile.Name())
	defer dbFile.Close()

	inflate := data.InflateCompressedFileInput{
		Task: in.DecompressTask,
		In:   tempDownload,
		Out:  dbFile,
	}

	if err := data.InflateCompressedFile(ctx, inflate); err != nil {
		return err
	}
	if _, err := dbFile.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Dir(in.FilePath)); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(in.FilePath), 0766); err != nil {
			return err
		}
	}

	final, err := os.OpenFile(in.FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer final.Close()

	_, err = io.Copy(final, dbFile)
	return err
}

// LastUpdatedAt checks the time that the database was last updated.
func (db *CardDB) LastUpdatedAt(ctx context.Context) (time.Time, error) {
	// 2023-02-28
	query := statement.Select("date").From("meta").Limit(1)

	var lastUpdatedAt time.Time
	if err := db.QueryRowStatement(ctx, query).Scan(&lastUpdatedAt); err != nil {
		return time.Time{}, err
	}

	// The database is generated in eastern, starting at noon and taking some number of hours
	loc, err := time.LoadLocation("EST")
	if err != nil {
		return time.Time{}, err
	}

	// We add 22 hours because the db file isn't updated till the evening every day
	return time.Date(lastUpdatedAt.Year(), lastUpdatedAt.Month(), lastUpdatedAt.Day(), 20, 0, 0, 0, loc), nil
}
