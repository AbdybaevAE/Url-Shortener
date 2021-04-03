package algo

import (
	"strings"

	typ_err "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	repo "github.com/abdybaevae/url-shortener/pkg/repos/algo"
	num_srv "github.com/abdybaevae/url-shortener/pkg/services/number"
)

type Strategy string

const (
	BASE_62 Strategy = "base_62"
)

func strategyFromString(val string) (Strategy, error) {
	switch val {
	case string(BASE_62):
		return BASE_62, nil
	}
	return "", typ_err.AlgoNotFound
}

type service struct {
	strategy Strategy
	entity   *models.Algo
	algoRepo repo.AlgoRepo
	numSrv   num_srv.NumberService
}

func newService(
	algoRepo repo.AlgoRepo, numSrv num_srv.NumberService, entity *models.Algo) (AlgoService, error) {
	strategy, err := strategyFromString(entity.Strategy)
	if err != nil {
		return nil, err
	}
	return &service{strategy: strategy, algoRepo: algoRepo, entity: entity, numSrv: numSrv}, nil
}
func (s *service) Entity() *models.Algo {
	return s.entity
}
func (s *service) GenerateKeys() ([]string, error) {
	ret := make([]string, 0)
	switch s.strategy {
	case BASE_62:
		numId, inc, dict := s.entity.NumberId, s.entity.IncrementValue, s.entity.Dict
		lastNum, err := s.numSrv.Increment(numId, inc)
		if err != nil {
			return nil, err
		}
		firstNum := lastNum - inc
		return getKeysByBaseStrategy(firstNum, lastNum, dict), nil
	// case [NEW_STRATEGY]
	default:
		return ret, typ_err.AlgoNotFound
	}
}
func isValidDictForBaseStrategy(dict string) bool {
	mp := make(map[rune]bool)
	for _, v := range dict {
		if _, ok := mp[v]; ok {
			return false
		}
		mp[v] = true
	}
	return true
}
func getKeysByBaseStrategy(from int, to int, dict string) []string {
	ret := make([]string, 0)
	if to < from || len(dict) <= 0 || !isValidDictForBaseStrategy(dict) {
		return ret
	}
	size := len(dict)
	for curr := from; curr <= to; curr++ {
		var sb strings.Builder
		num := curr
		for num > 0 {
			sb.WriteByte(dict[num%size])
			num /= size
		}
		ret = append(ret, sb.String())
	}
	return ret

}
