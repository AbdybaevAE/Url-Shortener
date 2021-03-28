package models

type Number struct {
	Id             int `database:"id"`
	LastNumber     int `database:"last_number"`
	IncrementValue int `database:"increment_value"`
}
