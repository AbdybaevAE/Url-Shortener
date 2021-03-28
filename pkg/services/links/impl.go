package links

import (
	"net/url"

	"github.com/abdybaevae/url-shortener/pkg/errors"
	"github.com/abdybaevae/url-shortener/pkg/models"
	linkRepo "github.com/abdybaevae/url-shortener/pkg/repos/links"
	"github.com/abdybaevae/url-shortener/pkg/services/keys"
)

type LinkServiceImpl struct {
	keyService keys.KeyService
	linkRepo   linkRepo.LinkRepo
}

func NewLinkService(keyService keys.KeyService, linkRepo linkRepo.LinkRepo) LinkService {
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
