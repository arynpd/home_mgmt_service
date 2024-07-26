package service

import "github.com/arynpd/home-mgmt-service/db"

type Service struct {
	db *db.Db
}

func (s *Service) Init() error {
	dbConn := &db.Db{}
	err := dbConn.Init()
	if err != nil {
		return err
	}

	s.db = dbConn
	return nil
}

func (s *Service) Close() {
	s.db.Close()
}
