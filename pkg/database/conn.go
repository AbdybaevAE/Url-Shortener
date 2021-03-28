package database

import "github.com/jmoiron/sqlx"

func Conn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://cifer@localhost:5432/url_shortener?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
