package number

import (
	repo "github.com/abdybaevae/url-shortener/pkg/repos/number"
)

type service struct {
	repo repo.NumberRepo
}

func New(numRepo repo.NumberRepo) NumberService {
	return &service{repo: numRepo}
}

func (s *service) Increment(numberId int, byValue int) (int, error) {
	return s.repo.Increment(numberId, byValue)
}
