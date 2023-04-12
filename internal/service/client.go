package service

import (
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
)
type ClientService struct {
	rep repository.Service
}

func NewClientService(rep repository.Service) *ClientService {
	return &ClientService{rep: rep}
}


func(c *ClientService) GetUsers(login string, password string) (v1.User, error){
	return c.rep.GetUsers(login, password)
}
func(c *ClientService) PostUsers() error{
	return c.rep.PostUsers()
}
func(c *ClientService) UpdateUsers(id int, login string, password string) error{
	return c.rep.UpdateUsers(id, login, password)
}
func(c *ClientService) DeleteUsers(id int) error{
	return c.rep.DeleteUsers(id)
}

func(c *ClientService) GetSites() error{
	return c.rep.GetSites()
}
func(c *ClientService) PostSites() error{
	return c.rep.PostSites()
}
func(c *ClientService) UpdateSites() error{
	return c.rep.UpdateSites()
}
func(c *ClientService) DeleteSites() error{
	return c.rep.DeleteSites()
}

func(c *ClientService) GetMainText() error{
	return c.rep.GetMainText()
}
func(c *ClientService) PostMainText() error{
	return c.rep.PostMainText()
}
func(c *ClientService) UpdateMainText() error{
	return c.rep.UpdateMainText()
}
func(c *ClientService) DeleteMainText() error{
	return c.rep.DeleteMainText()
}

func(c *ClientService) GetPageSites() error{
	return c.rep.GetPageSites()
}
func(c *ClientService) PostPageSites() error{
	return c.rep.PostPageSites()
}
func(c *ClientService) UpdatePageSites() error{
	return c.rep.UpdatePageSites()
}
func(c *ClientService) DeletePageSites() error{
	return c.rep.DeletePageSites()
}