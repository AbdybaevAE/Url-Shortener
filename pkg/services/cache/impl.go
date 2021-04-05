package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/abdybaevae/url-shortener/pkg/errors/typed"
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
	exp := time.Until(link.ExpiredAt)
	if err := i.client.Set(
		context.Background(),
		i.buildLinkExpirePath(link.Key),
		link.ExpiredAt, exp).Err(); err != nil {
		return err
	}
	if err := i.client.Set(
		context.Background(),
		i.buildLinkValuePath(link.Key),
		link.Link,
		exp).Err(); err != nil {
		return err
	}
	return nil
}

// const rTimeLay string = "2023-04-19T19:44:58"

func (i *impl) GetLinkByKey(key string) (*models.Link, error) {
	expVal, err := i.client.Get(context.Background(), i.buildLinkExpirePath(key)).Result()
	if err == redis.Nil {
		return nil, typed.KeyNotFound
	}
	if err != nil {
		return nil, err
	}
	linkValue, err := i.client.Get(context.Background(), i.buildLinkValuePath(key)).Result()
	if err == redis.Nil {
		return nil, typed.KeyNotFound
	}
	if err != nil {
		return nil, err
	}
	expiredAt, err := time.Parse(time.RFC3339, expVal)
	if err != nil {
		return nil, err
	}
	return &models.Link{Link: linkValue, ExpiredAt: expiredAt}, nil
}
func (i *impl) RemoveLink(link *models.Link) error {
	return i.client.Del(context.Background(),
		i.buildLinkExpirePath(link.Key),
		i.buildLinkValuePath(link.Key)).Err()
}
func (i *impl) buildLinkValuePath(linkKey string) string {
	return fmt.Sprintf("links:%s:link_value", linkKey)
}
func (i *impl) buildLinkExpirePath(linkKey string) string {
	return fmt.Sprintf("links:%s:expired_at", linkKey)
}
