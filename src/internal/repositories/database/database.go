package database

import "github.com/maddiesch/collector/internal/db"

type Database struct {
	conn *db.Conn
}

func New(conn *db.Conn) *Database {
	return &Database{
		conn: conn,
	}
}

func (d *Database) Close() error {
	return d.conn.Close()
}
