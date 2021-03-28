package link

import "github.com/abdybaevae/url-shortener/pkg/models"

type repo struct{}

func New() LinkRepo {
	return &repo{}
}
func (r *repo) Save(link *models.Link) error {
	return nil
}
func (r *repo) Get(key string) (*models.Link, error) {
	return nil, nil
}
