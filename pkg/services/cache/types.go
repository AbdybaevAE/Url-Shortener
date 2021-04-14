package cache

import (
	"github.com/abdybaevae/url-shortener/pkg/models"
	token_srv "github.com/abdybaevae/url-shortener/pkg/services/token"
)

type CacheService interface {
	SetLink(link *models.Link) (err error)
	GetLinkByKey(key string) (link *models.Link, err error)
	RemoveLink(link *models.Link) (err error)
	SetToken(user *models.User, token *token_srv.TokenDetails) (err error)
	HasTokenUuid(uuid string) (has bool, err error)
	DeleteTokenUuid(uuid string)
}
