// Package raptor is a Sqlite3 interface
package raptor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

const (
	DriverName = "sqlite3"
)

var (
	connID uint64
)

// New opens a new database connection
func New(source string) (*Conn, error) {
	db, err := sql.Open(DriverName, source)
	if err != nil {
		return nil, err
	}

	return &Conn{
		db:  db,
		log: zap.NewNop(),
		id:  atomic.AddUint64(&connID, 1),
	}, nil
}

type Conn struct {
	mu  sync.RWMutex // Config mutex
	id  uint64       // Connection id
	sp  uint64       // Savepoint id
	db  *sql.DB      // Underlying database connection
	log *zap.Logger  // Logger instance
}

// Close the database connection and perform any necessary cleanup
//
// Once close is called, new queries will be rejected.
// Close will block until all outstanding queries have completed.
func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.log.Sync()

	return c.db.Close()
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (c *Conn) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

// SetLogger assigns a logger instance to the connection. If you don't want logging, use `zap.NewNop()` (this is also the default)
func (c *Conn) SetLogger(log *zap.Logger) {
	c.mu.Lock()
	c.log = log
	c.mu.Unlock()
}

// A Result summarizes an executed SQL command.
type Result sql.Result

// Rows is the result of a query. See sql.Rows for more information.
type Rows sql.Rows

// Row is the result of calling QueryRow to select a single row.
type Row sql.Row

// Executor defines an interface for executing queries that don't return rows.
type Executor interface {
	Exec(context.Context, string, ...any) (Result, error)
}

// Exec perform a query on the database. It will not return any rows. e.g. insert or delete
func (c *Conn) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	return c.exec(ctx, query, args...)
}

func (c *Conn) exec(ctx context.Context, query string, args ...any) (Result, error) {
	c.mu.RLock()
	c.log.Debug("Exec", zap.String("query", query))
	c.mu.RUnlock()

	r, err := c.db.ExecContext(ctx, query, args...)

	return Result(r), err
}

// Querier defines an interface for executing queries that return rows from the database.
type Querier interface {
	Query(context.Context, string, ...any) (*Rows, error)
	QueryRow(context.Context, string, ...any) *Row
}

func (c *Conn) Query(ctx context.Context, query string, args ...any) (*Rows, error) {
	return c.query(ctx, query, args...)
}

func (c *Conn) query(ctx context.Context, query string, args ...any) (*Rows, error) {
	c.mu.RLock()
	c.log.Debug("Query", zap.String("query", query))
	c.mu.RUnlock()

	r, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return (*Rows)(r), nil
}

func (c *Conn) QueryRow(ctx context.Context, query string, args ...any) *Row {
	return c.queryRow(ctx, query, args...)
}

func (c *Conn) queryRow(ctx context.Context, query string, args ...any) *Row {
	c.mu.RLock()
	c.log.Debug("QueryRow", zap.String("query", query))
	c.mu.RUnlock()

	r := c.db.QueryRowContext(ctx, query, args...)

	return (*Row)(r)
}

func (c *Conn) newSavepointName() string {
	return fmt.Sprintf("tx_%d_%d", c.id, atomic.AddUint64(&c.sp, 1))
}

// TxRollbackError is returned when a transaction is rolled back and the rollback also returns an error.
type TxRollbackError struct {
	Underlying error
	Rollback   error
}

func (e *TxRollbackError) Error() string {
	return fmt.Sprintf("rollback error: %s; rollback error: %s", e.Underlying, e.Rollback)
}

// TxBroker defines an interface for performing a transaction.
type TxBroker interface {
	Transact(context.Context, func(DB) error) error
}

// DB defines a standard set of interfaces that allow CRUD operations on a database.
type DB interface {
	Executor
	Querier
	TxBroker
}

var _ DB = (*Conn)(nil)
var _ DB = (*txConn)(nil)

func (c *Conn) Transact(ctx context.Context, fn func(DB) error) error {
	return c.transact(ctx, fn)
}

func (c *Conn) transact(ctx context.Context, fn func(DB) error) error {
	savepoint := c.newSavepointName()

	txConn := &txConn{
		conn:  c,
		name:  savepoint,
		state: txStateInit,
	}

	if err := txConn.begin(ctx); err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			txConn.rollback(ctx)
			panic(p)
		}
	}()

	if err := fn(txConn); err != nil {
		if rErr := txConn.rollback(ctx); err != nil {
			return &TxRollbackError{Underlying: err, Rollback: rErr}
		}
		return err
	}

	return txConn.commit(ctx)
}

const (
	txStateInit uint8 = iota
	txStateRunning
	txStateCommitted
	txStateRollbacked
)

type txConn struct {
	conn  *Conn
	mu    sync.Mutex
	name  string
	state uint8
}

var (
	ErrTransactionAlreadyStarted = errors.New("transaction already started")
)

func (t *txConn) begin(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.state != txStateInit {
		return ErrTransactionAlreadyStarted
	}

	_, err := t.conn.exec(ctx, "SAVEPOINT "+t.name+";")
	if err == nil {
		t.state = 1
	}

	return err
}

func (t *txConn) rollback(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.state != txStateRunning {
		return nil
	}

	_, err := t.conn.exec(ctx, "ROLLBACK TRANSACTION TO SAVEPOINT "+t.name+";")
	if err == nil {
		t.state = txStateRollbacked
	}

	return nil
}

func (t *txConn) commit(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.state != txStateRunning {
		return nil
	}

	_, err := t.conn.exec(ctx, "RELEASE SAVEPOINT "+t.name+";")
	if err == nil {
		t.state = txStateCommitted
	}

	return err
}

func (t *txConn) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	return t.conn.exec(ctx, query, args...)
}

func (t *txConn) Transact(ctx context.Context, fn func(DB) error) error {
	return t.conn.transact(ctx, fn)
}

func (t *txConn) Query(ctx context.Context, query string, args ...any) (*Rows, error) {
	return t.conn.query(ctx, query, args...)
}

func (t *txConn) QueryRow(ctx context.Context, query string, args ...any) *Row {
	return t.conn.queryRow(ctx, query, args...)
}
