package service

import "github.com/arynpd/home-mgmt-service/db"

type Service struct {
	db *db.Db
}

func (s *Service) Init(connString string) error {
	dbConn := &db.Db{}
	err := dbConn.Init(connString)
	if err != nil {
		return err
	}

	s.db = dbConn
	return nil
}

func (s *Service) Close() {
	s.db.Close()
}
