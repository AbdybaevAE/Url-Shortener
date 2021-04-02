package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Conn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://cifer@localhost:5432/url_shortener?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
