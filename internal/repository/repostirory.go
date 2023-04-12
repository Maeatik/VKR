package repository

import (
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/pgsql"
)
type Authorization interface {
	CreateUser(user v1.User) (int, error)
	GetUser(name, password string) (v1.User, error)
}

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

type Repository struct {
	Authorization
	Service

}

func NewRepository(db *pgsql.Postgres) *Repository {
	return &Repository{
		Authorization: 	NewAuthPostgres(db),
		Service: NewClientService(db),
	}
}