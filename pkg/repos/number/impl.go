package number

import (
	typed_errors "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) NumberRepo {
	return &repo{db: db}
}

var incrementQuery = `
	update numbers 
	set 
		number_value = number_value + $1
		where number_id = $2
	returning number_value
`

func (r *repo) Increment(numberId int, byValue int) (int, error) {
	if byValue <= 0 {
		return 0, typed_errors.InvalidIncrementNumberValue
	}
	var lastNum int
	if err := r.db.QueryRow(incrementQuery, byValue, numberId).Scan(&lastNum); err != nil {
		return 0, err
	}
	return lastNum, nil
}
