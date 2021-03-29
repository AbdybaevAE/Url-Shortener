package models

type Key struct {
	Id     int    `db:"key_id"`
	Value  string `db:"key_value"`
	AlgoId int    `db:"algo_id"`
}
