package token

import "github.com/abdybaevae/url-shortener/pkg/models"

type TokenPair struct {
	Access  string
	Refresh string
}

type TokenService interface {
	GenerateToken(user *models.User) (pair *TokenDetails, err error)
	RefreshToken(refreshToken string) (pair *TokenDetails, err error)
}
