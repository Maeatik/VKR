package repository

import (
	"context"
	"crypto/sha1"
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/pgsql"
	"encoding/base64"
	"fmt"
	// "github.com/jmoiron/sqlx"
)

var salt = "drinkingeveryday1mdrinking"

type ClientServicePostgres struct {
	db *pgsql.Postgres
}

func NewClientService(db *pgsql.Postgres) *ClientServicePostgres  {
	return &ClientServicePostgres{db: db}
}

func GeneratePasswordHash(password string) string  {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hash.Sum([]byte(salt)))
	return fmt.Sprintf("%s", sha)
}

func(c *ClientServicePostgres) GetUsers(login string, password string) (v1.User, error){
	hashPassword := GeneratePasswordHash(password)
	fmt.Println(hashPassword)
	ctx := context.Background()
	var user v1.User
 	row := c.db.QueryRow(ctx, `SELECT login, password FROM "Users" WHERE password = $1`, hashPassword)
	if err := row.Scan(&user.Name, &user.Password); err != nil{
		fmt.Println(user.Id)
		return v1.User{}, err
	}

	return user, nil
}
//Не надо
func(c *ClientServicePostgres) PostUsers() error{
	return nil
}

func(c *ClientServicePostgres) UpdateUsers(id int, login string, password string) error{
	hashPassword := GeneratePasswordHash(password)

	ctx := context.Background()
	var user v1.User
	fmt.Println(login)
 	_, err := c.db.Exec(ctx, `UPDATE "Users" SET login = $1, password = $2 WHERE id=$3`, login, hashPassword, id)
	if err != nil{
		fmt.Println(user.Id)
		return err
	}

	return nil
}
func(c *ClientServicePostgres) DeleteUsers(id int) error{
	ctx := context.Background()

	_, err := c.db.Exec(ctx, `DELETE FROM "Users" where id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func(c *ClientServicePostgres) GetSites() error{
	return nil
}
func(c *ClientServicePostgres) PostSites() error{
	return nil
}
func(c *ClientServicePostgres) UpdateSites() error{
	return nil
}
func(c *ClientServicePostgres) DeleteSites() error{
	return nil
}

func(c *ClientServicePostgres) GetMainText() error{
	return nil
}
func(c *ClientServicePostgres) PostMainText() error{
	return nil
}
func(c *ClientServicePostgres) UpdateMainText() error{
	return nil
}
func(c *ClientServicePostgres) DeleteMainText() error{
	return nil
}

func(c *ClientServicePostgres) GetPageSites() error{
	return nil
}
func(c *ClientServicePostgres) PostPageSites() error{
	return nil
}
func(c *ClientServicePostgres) UpdatePageSites() error{
	return nil
}
func(c *ClientServicePostgres) DeletePageSites() error{
	return nil
}