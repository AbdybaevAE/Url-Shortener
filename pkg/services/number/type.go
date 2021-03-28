package number

import "github.com/abdybaevae/url-shortener/pkg/models"

// Number service that resonsible for increment number concurrently safe
// Once it was created, it's impossible to edit entity, better solution would be to create another one entity
type NumberService interface {
	// Create entity
	Create(number *models.Number) (id int, err error)
	// Increment entity value byValue and return last number result
	Increment(id int, byValue int) (lastNumber int, err error)
}
