package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Pool *pgxpool.Pool
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
