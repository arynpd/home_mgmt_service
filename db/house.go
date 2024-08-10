package db

import (
	"context"
	"fmt"

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
	stmt := `insert into home_schema.house (street, city, state, zip) 
			values ($1, $2, $3, $4) 
			returning id`

	err := db.connType.QueryRow(context.Background(), stmt, h.Street, h.City, h.State, h.Zip).Scan(&h.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) UpdateHouse(h *House) error {
	stmt := `update home_schema.house
			set street = $1, city = $2, state = $3, zip = $4
			where id = $5`

	tag, err := db.connType.Exec(context.Background(), stmt, h.Street, h.City, h.State, h.Zip, h.Id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Could not find house with id: %d", h.Id)
	}

	return nil
}

func (db *Db) GetHouseById(h *House) error {
	stmt := `select street, city, state, zip 
			from home_schema.house
			where id = $1`
	err := db.connType.QueryRow(context.Background(), stmt, h.Id).Scan(&h.Street, &h.City, &h.State, &h.Zip)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("Could not find house with id: %d", h.Id)
		}
		return err
	}
	return nil
}

func (db *Db) DeleteHouse(h *House) error {
	stmt := `delete
			from home_schema.house
			where id = $1`
	tag, err := db.connType.Exec(context.Background(), stmt, h.Id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Could not find house with id: %d", h.Id)
	}

	return nil
}
