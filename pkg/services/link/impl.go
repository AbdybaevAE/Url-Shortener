package links

import (
	"net/url"
	"time"

	http_errors "github.com/abdybaevae/url-shortener/pkg/errors/http"
	"github.com/abdybaevae/url-shortener/pkg/models"
	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	key_service "github.com/abdybaevae/url-shortener/pkg/services/key"
)

type LinkServiceImpl struct {
	keyService key_service.KeyService
	linkRepo   link_repo.LinkRepo
}

func New(linkRepo link_repo.LinkRepo, keyService key_service.KeyService) LinkService {

	return &LinkServiceImpl{linkRepo: linkRepo, keyService: keyService}
}

const DefaultExpireTime = time.Hour * 24 * 31 * 12 * 2

func (s *LinkServiceImpl) ShortenLink(longLink string) (string, error) {
	if longLink == "" {
		return "", http_errors.InvalidLink
	}
	_, err := url.ParseRequestURI(longLink)
	if err != nil {
		return "", http_errors.InvalidLink
	}
	key, err := s.keyService.Get()
	if err != nil {
		return "", err
	}
	link := &models.Link{
		Key:       key,
		Link:      longLink,
		ExpiredAt: time.Now().Add(DefaultExpireTime),
	}
	if err := s.linkRepo.Save(link); err != nil {
		return "", err
	}

	return key, nil
}
func (s *LinkServiceImpl) GetLink(key string) (string, error) {
	if key == "" {
		return "", http_errors.InvalidLinkKey
	}
	linkEntity, err := s.linkRepo.GetByKey(key)
	if err != nil {
		return "", err
	}
	return linkEntity.Link, nil
}
func (s *LinkServiceImpl) VisitByKey(key string) error {
	if key == "" {
		return http_errors.InvalidLinkKey
	}
	return s.linkRepo.VisitByKey(key)
}
