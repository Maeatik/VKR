package service

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/internal/repository"
	"diploma/internal/usecase"
	"fmt"
	"time"
)

type ClientService struct {
	rep repository.Service
}

func NewClientService(rep repository.Service) *ClientService {
	return &ClientService{rep: rep}
}

func (c *ClientService) GetUsers(ctx context.Context, id int) (v1.User, error) {
	return c.rep.GetUsers(ctx, id)
}

func (c *ClientService) UpdateUsers(ctx context.Context, id int, login string, password string) error {
	return c.rep.UpdateUsers(ctx, id, login, password)
}
func (c *ClientService) DeleteUsers(ctx context.Context, id int) error {
	if err := c.rep.DeleteAllTexts(ctx, id); err != nil {
		return err
	}

	if err := c.rep.DeleteAllSites(ctx, id); err != nil {
		return err
	}

	return c.rep.DeleteUsers(ctx, id)
}

func (c *ClientService) GetSite(ctx context.Context, userID int, id int) (v1.Site, error) {
	return c.rep.GetSite(ctx, userID, id)
}
func (c *ClientService) GetListSites(ctx context.Context, userID int) ([]v1.Site, error) {
	return c.rep.GetListSites(ctx, userID)
}
func (c *ClientService) PostSite(ctx context.Context, userID int, url string, tag string) (int, error) {
	return c.rep.PostSite(ctx, userID, url, tag)
}
func (c *ClientService) DeleteSite(ctx context.Context, userID int, id int) error {
	fmt.Println(userID, id)
	if err := c.rep.DeleteAllSiteTexts(ctx, userID, id); err != nil {
		return err
	}
	fmt.Println(2)
	return c.rep.DeleteSite(ctx, userID, id)
}

func (c *ClientService) GetMainText(ctx context.Context, userID int, id int, siteID int) (v1.MainText, error) {
	return c.rep.GetMainText(ctx, userID, id, siteID)
}
func (c *ClientService) PostMainText(ctx context.Context, userID int, siteID int, text string) error {
	date := time.Now()

	return c.rep.PostMainText(ctx, userID, siteID, date, text)
}
func (c *ClientService) UpdateMainText(ctx context.Context, userID int, id int, siteID int, text string) error {
	date := time.Now()

	return c.rep.UpdateMainText(ctx, userID, id, siteID, date, text)
}

func (c *ClientService) ParseSite(ctx context.Context, userID int, url, tag string) error {
	text, err := usecase.GetTextRelatedToTag(url, "\"p\"", tag)
	if err != nil{
		return err
	}

	siteID, err := c.rep.PostSite(ctx, userID, url, tag)
	if err != nil{
		return err
	}

	date := time.Now()

	return c.rep.PostMainText(ctx, userID, siteID, date, text)
}
