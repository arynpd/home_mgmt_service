package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type House struct {
	Id     int    `json:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    int    `json:"zip"`
}

func (db *Db) CreateHouse(h *House) error {
	stmt := `insert into home_schema.house (street, city, state, zip) values ($1, $2, $3, $4) returning id`
	tx, err := db.Pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), stmt, h.Street, h.City, h.State, h.Zip).Scan(&h.Id)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil

}
