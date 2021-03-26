package links

import (
	"github.com/abdybaevae/url-shortener/internal/services/keys"
	"github.com/abdybaevae/url-shortener/pkg/errors"
)

type LinkServiceImpl struct {
	keyService keys.KeyService
}

func NewLinkService(keyService keys.KeyService) LinkService {
	return &LinkServiceImpl{keyService: keyService}
}
func (s *LinkServiceImpl) Shorten(longLink string) (shortLink string, err error) {
	if shortLink == "" {
		return "", errors.InvalidLink
	}

	return "", nil
}
func (s *LinkServiceImpl) GetOriginalFromShorten(shortLink string) (longLink string, err error) {
	if shortLink == "" {
		return "", errors.InvalidLinkKey
	}
	return "", nil
}
