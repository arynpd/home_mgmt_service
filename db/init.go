package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Pool *pgxpool.Pool
}

type Transaction struct {
	transaction pgx.Tx
}

func (db *Db) Init() error {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Connected to database!'").Scan(&greeting)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", greeting)

	db.Pool = pool
	return nil
}

func (db *Db) Close() {
	db.Pool.Close()
}

func (db *Db) BeginTransaction() (*Transaction, error) {
	t := &Transaction{}
	tx, err := db.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	t.transaction = tx
	return t, nil

}

func (t *Transaction) Commit() error {
	return t.transaction.Commit(context.Background())
}

func (t *Transaction) Rollback() error {
	return t.transaction.Rollback(context.Background())
}
