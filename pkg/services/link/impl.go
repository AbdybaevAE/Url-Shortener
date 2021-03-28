package links

import (
	"net/url"

	"github.com/abdybaevae/url-shortener/pkg/errors"
	"github.com/abdybaevae/url-shortener/pkg/models"
	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	key_service "github.com/abdybaevae/url-shortener/pkg/services/key"
)

type LinkServiceImpl struct {
	keyService key_service.KeyService
	linkRepo   link_repo.LinkRepo
}

func New(linkRepo link_repo.LinkRepo, keyService key_service.KeyService) LinkService {
	return &LinkServiceImpl{keyService: keyService}
}
func (s *LinkServiceImpl) Shorten(longLink string) (string, error) {
	if longLink == "" {
		return "", errors.InvalidLink
	}
	_, err := url.ParseRequestURI(longLink)
	if err != nil {
		return "", errors.InvalidLink
	}
	key, err := s.keyService.Get()
	if err != nil {
		return "", nil
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
func (s *LinkServiceImpl) GetOriginalFromShorten(shortLink string) (longLink string, err error) {
	if shortLink == "" {
		return "", errors.InvalidLinkKey
	}
	return "", nil
}
