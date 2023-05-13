package service

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
)

type Service interface {
	GetUsers(ctx context.Context, id int) (v1.User, error)
	UpdateUsers(ctx context.Context, id int, login string, password string) error
	DeleteUsers(ctx context.Context, id int) error

	GetSite(ctx context.Context, userID int, id int) (v1.Site, error)
	GetListSites(ctx context.Context, userID int) ([]v1.Site, error)
	PostSite(ctx context.Context, userID int, url string, tag string) error
	DeleteSite(ctx context.Context, userID int, id int) error

	GetMainText(ctx context.Context, userID int, id int, siteID int) (v1.MainText, error)
	PostMainText(ctx context.Context, userID int, siteID int, text string) error
	UpdateMainText(ctx context.Context, userID int, id int, siteID int, text string) error
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

func NewServices(rep *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(rep.Authorization),
		Service:       NewClientService(rep.Service),
	}
}
