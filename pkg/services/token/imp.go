package token

import "github.com/abdybaevae/url-shortener/pkg/models"

type service struct{}

func New() TokenService {
	return &service{}
}
func (s *service) GenerateToken(user *models.User) (*TokenPair, error) {

	return nil, nil
}
func (s *service) RefreshToken(refreshToken string) (*TokenPair, error) {
	return nil, nil
}
func (s *service) DeleteToken(accessToken string) error {
	return nil
}
