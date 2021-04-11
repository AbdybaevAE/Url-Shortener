package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id      int    `db:"key_id"`
	Account string `db:"account"`
	Hash    string `db:"hash"`
}

func (u *User) GenHash(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Hash = string(hash)
	return nil
}
