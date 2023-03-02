package magic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/maddiesch/collector/internal/raptor"
	"github.com/maddiesch/collector/internal/raptor/statement"
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
