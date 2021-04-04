package link

import "github.com/abdybaevae/url-shortener/pkg/models"

type LinkRepo interface {
	Save(link *models.Link) (err error)
	GetByKey(key string) (link *models.Link, err error)
	VisitByKey(key string) (err error)
	Remove(link *models.Link) (err error)
}
