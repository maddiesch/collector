package database

import (
	"context"

	"github.com/maddiesch/collector/internal/db"
)

type Database struct {
	conn *db.Conn
}

func New(conn *db.Conn) *Database {
	return &Database{
		conn: conn,
	}
}

func (d *Database) Cleanup(ctx context.Context) error {
	return d.conn.Vacuum(ctx)
}

func (d *Database) Close() error {
	return d.conn.Close()
}
