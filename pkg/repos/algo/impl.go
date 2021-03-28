package algo

import (
	typed_errors "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) AlgoRepo {
	return &repo{db: db}
}

var getAlgoByQuery = "select * from algorithms where algo_name = $1"

func (r *repo) Get(algoName string) (*models.Algo, error) {
	var algo models.Algo
	if err := r.db.Get(&algo, getAlgoByQuery, getAlgoByQuery, algoName); err != nil {
		return nil, typed_errors.AlgoNotFound
	}
	return &algo, nil
}

// func (r *repo) someMEthod() {

// }
