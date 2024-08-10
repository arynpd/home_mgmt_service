package controller

import (
	"encoding/json"
	"net/http"

	"github.com/arynpd/home-mgmt-service/db"
)

func (c *Controller) CreateHouse(w http.ResponseWriter, r *http.Request) {
	var h db.House
	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		w.Write([]byte("decoder error"))
		return
	}

	err = c.service.CreateHouse(&h)
	if err != nil {
		w.Write([]byte("house creation error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h)
}
