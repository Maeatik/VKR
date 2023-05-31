package service

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Service interface {
	GetUser(ctx context.Context, id int) (v1.User, error)
	GetUsers(ctx context.Context, id int) (v1.User, error)
	UpdateUsers(ctx context.Context, id int, login string, password string) error
	ChangePassword(ctx context.Context, id int, password string, newPassword string) error
	DeleteUsers(ctx context.Context, id int, password string) error

	GetSite(ctx context.Context, userID int, id int) (v1.Site, error)
	GetListSites(ctx context.Context, userID int) ([]v1.Sites, error)
	PostSite(ctx context.Context, userID int, url string, tag string) (int, error)
	DeleteSite(ctx context.Context, userID int, id int) error

	GetMainText(ctx context.Context, userID int, id int, siteID int) (v1.MainText, error)
	PostMainText(ctx context.Context, userID int, siteID int, text string) error
	UpdateMainText(ctx context.Context, userID int, id int, siteID int, text string) error

	ParseSite(ctx context.Context, userID int, url, tag string) error
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
