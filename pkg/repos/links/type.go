package links

import "github.com/abdybaevae/url-shortener/pkg/models"

type LinkRepo interface {
	Save(link *models.Link) error
	Get(key string) (*models.Link, error)
}
