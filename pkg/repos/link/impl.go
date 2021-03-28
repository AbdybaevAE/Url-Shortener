package link

import (
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) LinkRepo {
	return &repo{
		db: db,
	}
}
func (r *repo) Save(link *models.Link) error {
	return nil
}
func (r *repo) Get(key string) (*models.Link, error) {
	return nil, nil
}
