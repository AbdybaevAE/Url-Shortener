package keys

import (
	"sync"

	typ_err "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	key_repo "github.com/abdybaevae/url-shortener/pkg/repos/key"
	"github.com/abdybaevae/url-shortener/pkg/services/algo"
	algo_srv "github.com/abdybaevae/url-shortener/pkg/services/algo"
)

type service struct {
	mu      *sync.Mutex
	keyRepo key_repo.KeyRepo
	algoSrv algo_srv.AlgoService
}

func New(keyRepo key_repo.KeyRepo, algoSrv algo.AlgoService) KeyService {
	return &service{
		keyRepo: keyRepo,
		algoSrv: algoSrv,
		mu:      &sync.Mutex{},
	}
}

// TODO Add circuit breaker
func (s *service) Get() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key, err := s.takeKey()
	if err == nil {
		return key, nil
	}
	if err != typ_err.NoKeys {
		return "", err
	}
	values, err := s.algoSrv.GenerateKeys()
	if err != nil {
		return "", err
	}
	keys := make([]models.Key, 0)
	for _, v := range values {
		keys = append(keys, models.Key{
			Value:  v,
			AlgoId: s.algoSrv.GetId(),
		})
	}
	if err := s.keyRepo.InsertMany(keys); err != nil {
		return "", err
	}
	return s.Get()
}
func (s *service) takeKey() (string, error) {
	return s.keyRepo.DeleteOne(s.algoSrv.GetId())
}
