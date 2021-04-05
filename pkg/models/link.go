package models

import "time"

type Link struct {
	Id        string    `db:"link_id"`
	Link      string    `db:"link_value"`
	Key       string    `db:"key_value"`
	CreatedAt time.Time `db:"created_at"`
	VisitedAt time.Time `db:"visited_at"`
	Visited   bool      `db:"visited"`
	ExpiredAt time.Time `db:"expired_at"`
}

func (l *Link) IsExpired() bool {
	return time.Now().After(l.ExpiredAt)
}
