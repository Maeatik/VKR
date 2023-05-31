package repository

import (
	"context"
	"crypto/sha1"
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/pgsql"
	"encoding/base64"
	"fmt"
	"time"
)

var salt = "thatsiscoursework"

type ClientServicePostgres struct {
	db *pgsql.Postgres
}

func NewClientService(db *pgsql.Postgres) *ClientServicePostgres {
	return &ClientServicePostgres{db: db}
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hash.Sum([]byte(salt)))
	return sha
}

func (c *ClientServicePostgres) GetUser(ctx context.Context, id int) (v1.User, error) {
	var user v1.User
	row := c.db.QueryRow(ctx, `SELECT login, password FROM "Users" WHERE id = $1`, id)
	if err := row.Scan(&user.Name, &user.Password); err != nil {
		fmt.Println(user.Id)
		return v1.User{}, err
	}

	return user, nil
}

func (c *ClientServicePostgres) GetUsers(ctx context.Context, id int) (v1.User, error) {
	var user v1.User
	row := c.db.QueryRow(ctx, `SELECT login, password FROM "Users" WHERE id = $1`, id)
	if err := row.Scan(&user.Name, &user.Password); err != nil {
		fmt.Println(user.Id)
		return v1.User{}, err
	}

	return user, nil
}

func (c *ClientServicePostgres) UpdateUsers(ctx context.Context, id int, login string, password string) error {
	var user v1.User

	_, err := c.db.Exec(ctx, `UPDATE "Users" SET login = $1, password = $2 WHERE id=$3`, login, password, id)
	if err != nil {
		fmt.Println(user.Id)
		return err
	}

	return nil
}
func (c *ClientServicePostgres) DeleteUsers(ctx context.Context, id int) error {
	_, err := c.db.Exec(ctx, `DELETE FROM "Users" where id=$1`, id)
	if err != nil {
		return fmt.Errorf("error while deleting user: %s", err.Error())
	}
	return nil
}

func (c *ClientServicePostgres) GetSite(ctx context.Context, userID int, id int) (v1.Site, error) {
	var site v1.Site

	row := c.db.QueryRow(ctx, `SELECT "url", tag FROM "Site" WHERE id = $1 AND "user_id(FK)" = $2`, id, userID)
	if err := row.Scan(&site.Url, &site.Tag); err != nil {
		return v1.Site{}, fmt.Errorf("error while getting site: %s", err.Error())
	}

	return site, nil
}
func (c *ClientServicePostgres) GetListSites(ctx context.Context, userID int) ([]v1.Sites, error) {
	var sites []v1.Sites

	query := `SELECT s.id as id_site, s.url, s.tag, mt.date, mt.id as id_text
				FROM "Site" s
				JOIN "Main_Text" mt ON s.id = mt."site_id(FK)"
				where mt."user_id(FK)" = $1 and s."user_id(FK)" = $1
				`
	rows, err := c.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of sites: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var site v1.Sites
		if err := rows.Scan(&site.IdSite, &site.Url, &site.Tag, &site.Date, &site.IdText); err != nil {
			return nil, fmt.Errorf("error while getting site: %s", err.Error())
		}

		sites = append(sites, site)
	}

	return sites, nil
}

func (c *ClientServicePostgres) PostSite(ctx context.Context, userID int, url string, tag string) (int, error) {
	var id int
	query := `INSERT INTO "Site" ("user_id(FK)", "url", tag) VALUES ($1, $2, $3)  REturning id`
	rows, err := c.db.Query(ctx, query, userID, url, tag)
	if err != nil {
		return 0, fmt.Errorf("error while insert new site result: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	fmt.Println(id)

	return id, nil
}

func (c *ClientServicePostgres) DeleteSite(ctx context.Context, userID int, id int) error {
	query := `DELETE FROM "Site" WHERE id=$1 AND "user_id(FK)"=$2`
	_, err := c.db.Exec(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("error while delete site: %s", err.Error())
	}
	return nil
}

func (c *ClientServicePostgres) GetMainText(ctx context.Context, userID int, id int, siteID int) (v1.MainText, error) {
	var text v1.MainText

	row := c.db.QueryRow(ctx, `SELECT "site_id(FK)", "date", maintext FROM "Main_Text" WHERE id = $1 AND "user_id(FK)" = $2 AND "site_id(FK)" = $3`, id, userID, siteID)
	if err := row.Scan(&text.SiteID, &text.Date, &text.Text); err != nil {
		return v1.MainText{}, fmt.Errorf("error while getting text: %s", err.Error())
	}

	return text, nil
}
func (c *ClientServicePostgres) PostMainText(ctx context.Context, userID int, siteID int, date time.Time, text string) error {
	query := `INSERT INTO "Main_Text" ("user_id(FK)", "site_id(FK)", "date", maintext) VALUES ($1, $2, $3, $4)`
	_, err := c.db.Exec(ctx, query, userID, siteID, date, text)
	if err != nil {
		return fmt.Errorf("error while insert new maintext result: %s", err.Error())
	}
	return nil
}
func (c *ClientServicePostgres) UpdateMainText(ctx context.Context, userID int, id int, siteID int, date time.Time, text string) error {
	fmt.Println(date, text, id, userID, siteID)
	query := `UPDATE "Main_Text" SET "date" = $1, maintext = $2 WHERE id=$3 AND "user_id(FK)"=$4 AND "site_id(FK)"=$5`
	_, err := c.db.Exec(ctx, query, date, text, id, userID, siteID)
	if err != nil {
		return fmt.Errorf("error while update text: %s", err.Error())
	}

	return nil
}

func (c *ClientServicePostgres) DeleteAllSites(ctx context.Context, userID int) error {
	query := `DELETE FROM "Site" WHERE "user_id(FK)"=$1`
	_, err := c.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("error while delete all sites: %s", err.Error())
	}
	return nil
}

func (c *ClientServicePostgres) DeleteAllTexts(ctx context.Context, userID int) error {
	query := `DELETE FROM "Main_Text" WHERE "user_id(FK)"=$1`
	_, err := c.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("error while delete all maintexts: %s", err.Error())
	}
	return nil
}

func (c *ClientServicePostgres) DeleteAllSiteTexts(ctx context.Context, userID int, id int) error {
	fmt.Println(id, userID)
	query := `DELETE FROM "Main_Text" WHERE "site_id(FK)"=$1 AND "user_id(FK)"=$2`
	_, err := c.db.Exec(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("error while delete all maintexts of tag: %s", err.Error())
	}
	return nil
}

func (c *ClientServicePostgres) ParseSite(ctx context.Context, userID int, url, tag string) error {
	return nil
}
