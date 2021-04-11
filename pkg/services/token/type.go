package token

import "github.com/abdybaevae/url-shortener/pkg/models"

type TokenPair struct {
	Access  string
	Refresh string
}

type TokenService interface {
	GenerateToken(user *models.User) (pair *TokenPair, err error)
	RefreshToken(refreshToken string) (pair *TokenPair, err error)
	DeleteToken(accessToken string) (err error)
}
