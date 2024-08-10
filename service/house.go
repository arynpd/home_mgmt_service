package service

import "github.com/arynpd/home-mgmt-service/db"

func (s *Service) CreateHouse(h *db.House) error {
	return s.db.WithTx(func() error {
		return s.db.CreateHouse(h)
	})

}
