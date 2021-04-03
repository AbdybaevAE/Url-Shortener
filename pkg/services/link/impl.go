package links

import (
	"net/url"

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
func (s *LinkServiceImpl) Shorten(longLink string) (string, error) {
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
		Key:  key,
		Link: longLink,
	}
	if err := s.linkRepo.Save(link); err != nil {
		return "", err
	}

	return key, nil
}
func (s *LinkServiceImpl) GetOriginalFromShorten(shortLink string) (string, error) {
	if shortLink == "" {
		return "", http_errors.InvalidLinkKey
	}
	return "", nil
}
