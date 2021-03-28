package dict

import "github.com/abdybaevae/url-shortener/pkg/models"

type DictService interface {
	Create(dict *models.Dict) error
	Increment(dict *models.Dict) (lastNum int, err error)
}
