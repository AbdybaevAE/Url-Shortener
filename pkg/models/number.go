package models

type Number struct {
	Id    int `db:"number_id"`
	Value int `db:"number_value"`
}
