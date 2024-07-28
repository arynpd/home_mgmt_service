package db_test

import (
	"fmt"
	"testing"

	"github.com/arynpd/home-mgmt-service/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateHouse(t *testing.T) {
	dbPool := setupDbPool()
	t.Run("Create house success", func(t *testing.T) {
		house := &db.House{
			Street: "6118 Bummy Ln",
			City:   "Springfiled",
			State:  "TX",
			Zip:    69420,
		}

		err := dbPool.CreateHouse(house)
		assert.NoError(t, err)
	})
}

func TestUpdateHouse(t *testing.T) {
	dbPool := setupDbPool()

	t.Run("Update house success", func(t *testing.T) {
		house := &db.House{
			Street: "6118 Bummy Ln",
			City:   "Springfiled",
			State:  "TX",
			Zip:    69420,
		}
		err := dbPool.CreateHouse(house)
		assert.NoError(t, err)

		house.State = "VA"
		err = dbPool.UpdateHouse(house)
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

		err := dbPool.UpdateHouse(house)
		assert.EqualError(t, err, fmt.Sprintf("Could not find house with id: %d", house.Id))
	})
}
