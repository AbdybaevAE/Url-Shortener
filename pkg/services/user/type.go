package user

import token_srv "github.com/abdybaevae/url-shortener/pkg/services/token"

type UserService interface {
	Register(account, password string) (err error)
	Login(account, password string) (p *token_srv.TokenPair, err error)
	RefreshToken(refreshToken string) (p *token_srv.TokenPair, err error)
	Logout(accessToken string) (err error)
}
