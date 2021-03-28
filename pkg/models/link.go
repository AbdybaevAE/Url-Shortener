package models

import "time"

type Link struct {
	Id        string
	Link      string
	Key       string
	CreatedAt time.Time
	VisitedAt time.Time
	ExpiredAt time.Time
	IsActive  bool
}
