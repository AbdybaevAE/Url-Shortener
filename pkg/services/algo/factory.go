package algo

import (
	"log"
	"sync"

	http_errors "github.com/abdybaevae/url-shortener/pkg/errors/http"
	repo "github.com/abdybaevae/url-shortener/pkg/repos/algo"
	num_srv "github.com/abdybaevae/url-shortener/pkg/services/number"
)

type factory struct {
	mu         *sync.Mutex
	repo       repo.AlgoRepo
	store      map[string]AlgoService
	numService num_srv.NumberService
}

func NewFactory(algoRepo repo.AlgoRepo, numService num_srv.NumberService) AlgoFactory {
	if algoRepo == nil || numService == nil {
		log.Fatalln("Cannot init algo factory")
	}
	return &factory{
		repo:       algoRepo,
		store:      make(map[string]AlgoService),
		numService: numService,
		mu:         &sync.Mutex{},
	}
}
func (f *factory) Get(strategy string) (AlgoService, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if val, ok := f.store[strategy]; ok {
		return val, nil
	}
	entity, err := f.repo.Get(strategy)
	if err != nil {
		return nil, err
	}
	algoService, err := newService(f.repo, f.numService, entity)
	if err != nil {
		return nil, http_errors.ServerInternal
	}
	f.store[strategy] = algoService
	return algoService, nil
}
