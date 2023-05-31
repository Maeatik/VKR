package repository

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/pgsql"
	"time"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type Authorization interface {
	CreateUser(user v1.User) (int, error)
	GetUser(name, password string) (v1.User, error)
}

type Service interface {
	GetUser(ctx context.Context, id int) (v1.User, error)
	GetUsers(ctx context.Context, id int) (v1.User, error)
	UpdateUsers(ctx context.Context, id int, login string, password string) error
	DeleteUsers(ctx context.Context, id int) error

	GetSite(ctx context.Context, userID int, id int) (v1.Site, error)
	GetListSites(ctx context.Context, userID int) ([]v1.Sites, error)
	PostSite(ctx context.Context, userID int, url string, tag string) (int, error)
	DeleteSite(ctx context.Context, userID int, id int) error

	GetMainText(ctx context.Context, userID int, id int, siteID int) (v1.MainText, error)
	PostMainText(ctx context.Context, userID int, siteID int, date time.Time, text string) error
	UpdateMainText(ctx context.Context, userID int, id int, siteID int, date time.Time, text string) error

	DeleteAllSiteTexts(ctx context.Context, userID int, id int) error
	DeleteAllTexts(ctx context.Context, userID int) error
	DeleteAllSites(ctx context.Context, userID int) error

	ParseSite(ctx context.Context, userID int, url, tag string) error
}

type Repository struct {
	Authorization
	Service
}

func NewRepository(db *pgsql.Postgres) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Service:       NewClientService(db),
	}
}
