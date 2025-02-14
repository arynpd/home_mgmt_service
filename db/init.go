package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type connType interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

type options interface {
	Close()
}

type Db struct {
	pool connType
}

func (db *Db) Init(connString string) error {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Connected to database!'").Scan(&greeting)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", greeting)

	db.pool = pool
	return nil
}

func (db *Db) Close() {
	if poolOptions, ok := db.pool.(options); ok {
		poolOptions.Close()
	}
}

func (db *Db) ExecFile(filePath string) error {
	c, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	sql := string(c)
	_, err = db.pool.Exec(context.Background(), sql)
	return err
}

func (db *Db) WithTx(txFunc func() error) error {
	tx, err := db.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	origConn := db.pool
	db.pool = tx

	defer func() {
		tx.Rollback(context.Background())
		db.pool = origConn
	}()

	err = txFunc()
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
