package algo

import "github.com/abdybaevae/url-shortener/pkg/models"

type AlgoRepo interface {
	// Create(algo *models.Algorithm) (err error)
	// GetAll() (algo *models.Algorithm, err error)
	// Get(name string) (algo *models.Algorithm, err error)
	Get(name string) (algo *models.Algo, err error)
}
