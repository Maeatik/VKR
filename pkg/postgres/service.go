package postgres

import (
	"diploma/config"
	"diploma/pkg/logger"
	"diploma/pkg/pgsql"
	"fmt"

	//"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	Db *sqlx.DB
}

func New(cfg config.Config, log *logger.Logger) (*pgsql.Postgres, error) {
	psqlInfo := "postgres://postgres:pass@localhost:5432/MortyGRAB?sslmode=disable"
	//psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	fmt.Println(psqlInfo)
	db, err := pgsql.New(psqlInfo)

	if err != nil {
		log.Info(err.Error())
	}

	//Возвращение указателя на новую БД
	return db, err
}

func (r *PostgresRepository) Close() {
	r.Db.Close()
}
