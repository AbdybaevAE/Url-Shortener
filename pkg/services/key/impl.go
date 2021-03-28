package keys

import (
	key_repo "github.com/abdybaevae/url-shortener/pkg/repos/key"
	algo_service "github.com/abdybaevae/url-shortener/pkg/services/algo"
)

type KeyServiceImpl struct{}

func New(keysRepo key_repo.KeyRepo, algosService algo_service.AlgoService) KeyService {
	return &KeyServiceImpl{}
}
func (s *KeyServiceImpl) Get() (string, error) {

	return "", nil
}
