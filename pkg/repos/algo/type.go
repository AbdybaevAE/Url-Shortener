package algo

import "github.com/abdybaevae/url-shortener/pkg/models"

type AlgoRepo interface {
	Get(name string) (algo *models.Algo, err error)
}
