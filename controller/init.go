package controller

import "github.com/arynpd/home-mgmt-service/service"

type Controller struct {
	service *service.Service
}

func (c *Controller) Init() error {
	s := &service.Service{}
	err := s.Init()
	if err != nil {
		return err
	}

	c.service = s
	return nil
}

func (c *Controller) Close() {
	c.service.Close()
}
