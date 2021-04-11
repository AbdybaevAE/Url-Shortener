package user

import (
	typederr "github.com/abdybaevae/url-shortener/pkg/errors/typed"
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

const getUserByAccountQuery = `
	select * from users 
	where account = $1
`

func (r *repo) GetByAccount(account string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Get(user, getUserByAccountQuery, account); err != nil {
		return nil, typederr.UserNotFound
	}
	return user, nil
}
