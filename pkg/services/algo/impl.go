package algo

import (
	repo "github.com/abdybaevae/url-shortener/pkg/repos/algo"
)

type service struct{}

const DefaultAlgorithmName = "base_62"

func New(algoRepo repo.AlgoRepo) AlgoService {
	return &service{}
}

func (s *service) GenerateKeys(name string) ([]string, error) {
	return nil, nil
}
func (s *service) EnsureAll() error {
	return nil
}
