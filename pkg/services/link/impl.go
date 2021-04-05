package links

import (
	"fmt"
	"net/url"
	"time"

	http_errors "github.com/abdybaevae/url-shortener/pkg/errors/http"
	"github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	cache_srv "github.com/abdybaevae/url-shortener/pkg/services/cache"
	key_srv "github.com/abdybaevae/url-shortener/pkg/services/key"
)

type LinkServiceImpl struct {
	linkRepo link_repo.LinkRepo
	keySrv   key_srv.KeyService
	cacheSrv cache_srv.CacheService
}

func New(linkRepo link_repo.LinkRepo, keySrv key_srv.KeyService, cacheSrv cache_srv.CacheService) LinkService {
	return &LinkServiceImpl{
		linkRepo: linkRepo,
		keySrv:   keySrv,
		cacheSrv: cacheSrv,
	}
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
	key, err := s.keySrv.Get()
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
	linkEntity, err := s.cacheSrv.GetLinkByKey(key)
	var notCached bool
	if err != nil {
		if err != typed.KeyNotFound {
			return "", http_errors.KeyNotFound
		}
		linkEntity, err = s.linkRepo.GetByKey(key)
		if err != nil {
			return "", err
		}
		notCached = true
	}
	if linkEntity.IsExpired() {
		s.archiveLinkSilent(linkEntity)
		return "", http_errors.KeyNotFound
	}
	if notCached {
		go s.cacheSrv.SetLink(linkEntity)
	}
	return linkEntity.Link, nil
}
func (s *LinkServiceImpl) archiveLinkSilent(link *models.Link) {
	go s.archiveLink(link)
}
func (s *LinkServiceImpl) archiveLink(link *models.Link) {
	fmt.Printf("Delete link %v \n", link)
	if err := s.linkRepo.Remove(link); err != nil {
		fmt.Printf("Error deleting link from db %v \n", link)
	}
	if err := s.cacheSrv.RemoveLink(link); err != nil {
		fmt.Printf("Error deleting link from cache %v \n", link)
	}
}
