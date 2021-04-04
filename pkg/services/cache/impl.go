package cache

import (
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/go-redis/redis/v8"
)

type impl struct {
	client *redis.Client
}

func New(client *redis.Client) CacheService {
	return &impl{
		client: client,
	}
}
func (i *impl) SetLink(link *models.Link) error {
	return nil
}
func (i *impl) GetLinkByKey(key string) (*models.Link, error) {

	return nil, nil
}
func (i *impl) RemoveLink(link *models.Link) error {
	return nil
}
func (*impl) VisitLink(link *models.Link) {

}
