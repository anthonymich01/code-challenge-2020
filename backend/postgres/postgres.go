package postgres

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func InitDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Database: os.Getenv("PGDATABASE"),
	})

	return db
}
