package db

import (
	"context"
	"fmt"
)

type House struct {
	Id     int    `json:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    int    `json:"zip"`
}

func (t *Transaction) CreateHouse(h *House) error {
	stmt := `insert into home_schema.house (street, city, state, zip) 
			values ($1, $2, $3, $4) 
			returning id`

	err := t.transaction.QueryRow(context.Background(), stmt, h.Street, h.City, h.State, h.Zip).Scan(&h.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) UpdateHouse(h *House) error {
	stmt := `update home_schema.house
			set street = $1, city = $2, state = $3, zip = $4
			where id = $5`

	tag, err := t.transaction.Exec(context.Background(), stmt, h.Street, h.City, h.State, h.Zip, h.Id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Could not find house with id: %d", h.Id)
	}

	return nil
}

func (t *Transaction) GetHouseById(h *House) error {
	stmt := `select street, city, state, zip 
			from home_schema.house
			where id = $1`
	err := t.transaction.QueryRow(context.Background(), stmt, h.Id).Scan(&h.Street, &h.City, &h.State, &h.Zip)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) DeleteHouse(h *House) error {
	stmt := `delete
			from home_schema.house
			where id = $1`
	tag, err := t.transaction.Exec(context.Background(), stmt, h.Id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Could not find house with id: %d", h.Id)
	}

	return nil
}
