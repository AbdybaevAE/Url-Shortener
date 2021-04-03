package database

import (
	"fmt"

	"github.com/abdybaevae/url-shortener/pkg/conf"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Conn(conf *conf.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.DbUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
