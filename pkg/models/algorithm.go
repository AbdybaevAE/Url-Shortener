package models

type Algorithm struct {
	Id       int    `database:"id"`
	Name     int    `database:"name"`
	NumberId int    `database:"number_id"`
	Metadata string `database:"metadata"`
}
