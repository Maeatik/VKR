package repository

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"diploma/pkg/pgsql"
	"fmt"

	// "errors"
	// "fmt"
)

type AuthPostgres struct {
	db *pgsql.Postgres
}

func NewAuthPostgres(db *pgsql.Postgres) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user v1.User) (int, error) {
	ctx := context.Background()
	
	var id int
	var count int
	var passwords []string
	roleQuery := `SELECT COUNT(*) FROM "Users"`
	rowQuery := a.db.QueryRow(ctx, roleQuery)

	if err := rowQuery.Scan(&count); err != nil {
		return 0, err
	}

	isUniq := true
	checkQuery := `SELECT password FROM "Users"`

	rows, err := a.db.Query(ctx, checkQuery)
	if err != nil{
		return 0, err
	}
	defer rows.Close()
	for rows.Next(){
		var password string

		err = rows.Scan(&password)
		if err != nil{
			return 0, err
		}
		passwords = append(passwords, password)
	}

	for _, password := range passwords {
		if password == user.Password {
			isUniq = false
			fmt.Println(isUniq)
		}
	}

	query := `INSERT INTO "Users" (login, password) VALUES ($1, $2) RETURNING id`
	if isUniq {
		if count == 0 {
			row := a.db.QueryRow(ctx, query, user.Name, user.Password)
			if err := row.Scan(&id); err != nil {
				return 0, err
			}
			return id, nil
		} else {
			row := a.db.QueryRow(ctx, query,  user.Name, user.Password)
			if err := row.Scan(&id); err != nil {
				return 0, err
			}
			return id, nil
		}
	} else {
		return 0, fmt.Errorf("password is busy")
	}
}

func (a *AuthPostgres) GetUser(login, password string) (v1.User, error) {
	ctx := context.Background()
	
	var user v1.User
	fmt.Println(login + " " + password)
	query := `SELECT id FROM "Users" WHERE login=$1 AND password=$2`
	row := a.db.QueryRow(ctx, query, login, password)
	if err := row.Scan(&user.Id); err != nil {
		return v1.User{}, err
	}

	return user, nil
}
