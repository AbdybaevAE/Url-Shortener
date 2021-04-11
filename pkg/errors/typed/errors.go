package typed

import "errors"

var (
	NoKeys                      = errors.New("No keys")
	KeyInsertError              = errors.New("Cannot create new keys in database")
	AlgoNotFound                = errors.New("Algorithm not found")
	NumberNotFound              = errors.New("Number not found")
	InvalidIncrementNumberValue = errors.New("Invalid increment value")
	LinkNotFound                = errors.New("Link not found")
	KeyNotFound                 = errors.New("Key not found")
	UserNotFound                = errors.New("User not found")
)
