package service

import (
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
)

type Service interface {
	GetUsers(login string, password string) (v1.User, error)
	PostUsers() error
	UpdateUsers(id int, login string, password string) error
	DeleteUsers(id int) error

	GetSites()  error
	PostSites() error
	UpdateSites() error
	DeleteSites()  error

	GetMainText()  error
	PostMainText() error
	UpdateMainText()  error
	DeleteMainText()  error

	GetPageSites() error
	PostPageSites() error
	UpdatePageSites()  error
	DeletePageSites() error
}

type Authorization interface {
	CreateUser(user v1.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Authorization
	Service
}

func NewServices(rep *repository.Repository) *Services  {
	return &Services{
		Authorization: NewAuthService(rep.Authorization),
		Service: NewClientService(rep.Service),
	}
}

