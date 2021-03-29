package models

type Algo struct {
	Id             int    `db:"algo_id"`
	Strategy       string `db:"algo_strategy"`
	NumberId       int    `db:"number_id"`
	IncrementValue int    `db:"increment_value"`
	Dict           string `db:"dict"`
}
