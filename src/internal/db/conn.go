package db

import (
	"bytes"
	"context"
	"crypto/md5"
	"database/sql"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/maddiesch/collector/internal/db/statement"
	"github.com/maddiesch/collector/internal/db/statement/conditional"
	"github.com/maddiesch/collector/internal/db/statement/generator"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type Conn struct {
	*sql.DB
}

func (c *Conn) Transaction(ctx context.Context, fn func(*Tx) error) error {
	tx, err := c.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(&Tx{tx}); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("transaction rollback failure: %v", err)
		}
		return err
	}

	return tx.Commit()
}

func (c *Conn) ExecStatement(ctx context.Context, statement generator.Generator) error {
	query, args, err := statement.Generate()
	if err != nil {
		return err
	}

	_, err = c.ExecContext(ctx, query, args...)

	return err
}

func (c *Conn) QueryStatement(ctx context.Context, statement generator.Generator) (*sql.Rows, error) {
	query, args, err := statement.Generate()
	if err != nil {
		return nil, err
	}

	return c.QueryContext(ctx, query, args...)
}

type ResultRow struct {
	*sql.Rows
}

func (c *Conn) EachRow(ctx context.Context, statement generator.Generator, fn func(*ResultRow) error) error {
	rows, err := c.QueryStatement(ctx, statement)
	if err != nil {
		return err
	}
	for rows.Next() {
		err = fn(&ResultRow{rows})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Conn) QueryStatementRow(ctx context.Context, statement generator.Generator) *sql.Row {
	query, args, err := statement.Generate()
	if err != nil {
		panic(err)
	}

	return c.QueryRowContext(ctx, query, args...)
}

type Tx struct {
	*sql.Tx
}

func (t *Tx) ExecStatement(ctx context.Context, statement generator.Generator) error {
	query, args, err := statement.Generate()
	if err != nil {
		return err
	}

	_, err = t.ExecContext(ctx, query, args...)

	return err
}

func (t *Tx) QueryStatement(ctx context.Context, statement generator.Generator) (*sql.Rows, error) {
	query, args, err := statement.Generate()
	if err != nil {
		return nil, err
	}

	return t.QueryContext(ctx, query, args...)
}

func (t *Tx) QueryStatementRow(ctx context.Context, statement generator.Generator) *sql.Row {
	query, args, err := statement.Generate()
	if err != nil {
		panic(err)
	}

	return t.QueryRowContext(ctx, query, args...)
}

type NewConnInput struct {
	FilePath string
}

func NewConn(ctx context.Context, in NewConnInput) (*Conn, error) {
	connStr := fmt.Sprintf("file:%s", in.FilePath)

	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	conn := &Conn{DB: db}

	if err := conn.performPendingMigration(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Conn) Close() error {
	return c.DB.Close()
}

//go:embed migrations/*.sql
var migrationFS embed.FS

func (c *Conn) performPendingMigration(ctx context.Context) error {
	if _, err := c.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS "Migration" ("Key" TEXT NOT NULL UNIQUE);`); err != nil {
		return errors.Wrap(err, "Create migration table failed")
	}

	paths, err := fs.Glob(migrationFS, "**/*.sql")
	if err != nil {
		return err
	}

	for _, path := range paths {
		content, err := fs.ReadFile(migrationFS, path)
		if err != nil {
			return err
		}

		if err := c.performMigration(ctx, createHashedIdentifier(path), content); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) performMigration(ctx context.Context, id string, content []byte) error {
	var exists bool

	err := c.QueryStatementRow(ctx,
		statement.Exists(statement.Select("1").From("Migration").Where(conditional.Equal("Key", id))),
	).Scan(&exists)

	if err != nil || exists {
		return err
	}

	commands := lo.Filter(lo.Map(bytes.Split(content, []byte{';'}), func(b []byte, _ int) string {
		return strings.TrimSpace(string(b))
	}), func(s string, _ int) bool {
		return s != ""
	})

	return c.Transaction(ctx, func(tx *Tx) error {
		for _, command := range commands {
			if _, err := tx.ExecContext(ctx, command); err != nil {
				return err
			}
		}

		return tx.ExecStatement(ctx, statement.Insert().Into("Migration").Value("Key", id))
	})
}

func createHashedIdentifier(v string) string {
	hash := md5.Sum([]byte(v))
	return hex.EncodeToString(hash[:])
}
