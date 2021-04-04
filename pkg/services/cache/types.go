package cache

import "github.com/abdybaevae/url-shortener/pkg/models"

type CacheService interface {
	SetLink(link *models.Link) (err error)
	GetLinkByKey(key string) (link *models.Link, err error)
	RemoveLink(link *models.Link) (err error)
}
