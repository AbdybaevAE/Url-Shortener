package typed

import "errors"

var NoKeys = errors.New("No keys")
var KeyInsertError = errors.New("Cannot create new keys in database")
var AlgoNotFound = errors.New("Algorithm not found")
var NumberNotFound = errors.New("Number not found")
var InvalidIncrementNumberValue = errors.New("Invalid increment value")
