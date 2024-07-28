package db_test

import (
	"fmt"
	"testing"

	"github.com/arynpd/home-mgmt-service/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateHouse(t *testing.T) {
	transaction := setupTransaction()
	defer transaction.Rollback()

	t.Run("Create house success", func(t *testing.T) {
		house := &db.House{
			Street: "6118 Bummy Ln",
			City:   "Springfiled",
			State:  "TX",
			Zip:    69420,
		}

		err := transaction.CreateHouse(house)
		assert.NoError(t, err)

	})

}

func TestUpdateHouse(t *testing.T) {
	transaction := setupTransaction()
	defer transaction.Rollback()

	t.Run("Update house success", func(t *testing.T) {
		house := &db.House{
			Street: "6118 Bummy Ln",
			City:   "Springfiled",
			State:  "TX",
			Zip:    69420,
		}
		err := transaction.CreateHouse(house)
		assert.NoError(t, err)

		house.State = "VA"
		err = transaction.UpdateHouse(house)
		assert.NoError(t, err)
	})

	t.Run("Update house not found fail", func(t *testing.T) {
		house := &db.House{
			Id:     500,
			Street: "Fail street",
			City:   "Fail city",
			State:  "VA",
			Zip:    21421,
		}

		err := transaction.UpdateHouse(house)
		assert.EqualError(t, err, fmt.Sprintf("Could not find house with id: %d", house.Id))
	})

}
