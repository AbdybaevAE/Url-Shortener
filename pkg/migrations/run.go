package migrations

import (
	algos_service "github.com/abdybaevae/url-shortener/pkg/services/algo"
)

type migrations struct {
	algoService algos_service.AlgoService
}
type migrationFunc = func(m *migrations) error

func Run(algoService algos_service.AlgoService) error {
	m := &migrations{algoService: algoService}
	all := []migrationFunc{
		ensureAllAlgorithms,
	}
	for _, f := range all {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil

}

func ensureAllAlgorithms(m *migrations) error {
	m.algoService.EnsureAll()
	return nil
}
