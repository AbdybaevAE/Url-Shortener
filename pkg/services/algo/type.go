package algo

import "github.com/abdybaevae/url-shortener/pkg/models"

type AlgoService interface {
	GenerateKeys() (values []string, err error)
	Entity() (algo *models.Algo)
}
type AlgoFactory interface {
	Get(algoName string) (AlgoService, error)
}
