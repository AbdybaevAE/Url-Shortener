package link

import (
	"time"

	http_err "github.com/abdybaevae/url-shortener/pkg/errors/http"
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

const insertNamedQuery = `
	insert into links 
		(link_value, key_value, expired_at)
	values 
		(:link_value, :key_value, :expired_at)
	`

func (r *repo) Save(link *models.Link) error {
	res, err := r.db.NamedExec(insertNamedQuery, link)
	if err != nil {
		return err
	}
	if _, err := res.RowsAffected(); err != nil {
		return err
	}
	return nil
}

const getByKeyQuery = `
	select * from links 
	where
		key_value = $1
`

func (r *repo) GetByKey(key string) (*models.Link, error) {
	link := &models.Link{}
	if err := r.db.Get(link, getByKeyQuery, key); err != nil {
		return nil, http_err.KeyNotFound
	}
	return link, nil
}

const visiteByKeyQuery = `
	update links
	set 
		visited_at = $1,
		visited = true
	where 
		key_value = $2
		
`

func (r *repo) VisitByKey(key string) error {
	if _, err := r.db.Exec(visiteByKeyQuery, time.Now(), key); err != nil {
		return err
	}
	return nil
}

const deleteLinkById = `
	delete from links
	where 
		link_id = $1
`

func (r *repo) Remove(link *models.Link) error {
	if _, err := r.db.Exec(deleteLinkById, link.Id); err != nil {
		return err
	}
	return nil
}
