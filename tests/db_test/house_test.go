package db_test

import (
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

	t.Run("Create house fail", func(t *testing.T) {
		house := &db.House{
			City:  "Springfiled",
			State: "TX",
			Zip:   69420,
		}
		err := dbPool.CreateHouse(house)
		assert.Error(t, err)
	})
}
