package user

import (
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepo {
	return &repo{
		db: db,
	}
}

const createUserQuery = ` 
	insert into users 
	(account, hash) values ($1, $2)
`

func (r *repo) Create(user *models.User) error {
	if _, err := r.db.Exec(createUserQuery, user.Account, user.Hash); err != nil {
		return err
	}
	return nil
}
