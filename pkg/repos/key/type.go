package key

import "github.com/abdybaevae/url-shortener/pkg/models"

type KeyRepo interface {
	DeleteOne(algoId int) (value string, err error)
	InsertMany(keys []models.Key) (err error)
}
