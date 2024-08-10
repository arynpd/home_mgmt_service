package service

import "github.com/arynpd/home-mgmt-service/db"

func (s *Service) CreateHouse(h *db.House) error {
	err := s.db.WithTx(func() error {
		if err := s.db.CreateHouse(h); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
