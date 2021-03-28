package models

type Key struct {
	Id     int
	Value  string `db:"value"`
	AlgoId int    `db:"algo_id"`
}
