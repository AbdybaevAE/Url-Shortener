package user

import "github.com/abdybaevae/url-shortener/pkg/models"

type UserRepo interface {
	Create(user *models.User) (err error)
}
