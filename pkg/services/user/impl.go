package user

import (
	http_err "github.com/abdybaevae/url-shortener/pkg/errors/http"
	"github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	usr_repo "github.com/abdybaevae/url-shortener/pkg/repos/user"
	token_srv "github.com/abdybaevae/url-shortener/pkg/services/token"
)

type service struct {
	repo     usr_repo.UserRepo
	tokenSrv token_srv.TokenService
}

func New(repo usr_repo.UserRepo, tokenSrv token_srv.TokenService) UserService {
	return &service{
		repo:     repo,
		tokenSrv: tokenSrv,
	}
}

// Register User by account and password
// TODO add better validation
func (s *service) Register(account, password string) error {
	if account == "" || password == "" {
		return http_err.InvalidUserData
	}
	u := &models.User{
		Account: account,
	}
	if err := u.GenHash(password); err != nil {
		return http_err.ServerInternal
	}
	if err := s.repo.Create(u); err != nil {
		return err
	}
	return nil
}
func (s *service) Login(account, password string) (*token_srv.TokenPair, error) {
	if account == "" || password == "" {
		return nil, http_err.InvalidUserData
	}
	user, err := s.repo.GetByAccount(account)
	if err != nil {
		return nil, err
	}
	if !user.IsValidPassword(password) {
		return nil, typed.UserNotFound
	}
	tokenPair, err := s.tokenSrv.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}
func (s *service) RefreshToken(refreshToken string) (p *token_srv.TokenPair, err error) {
	return nil, nil
}
func (s *service) Logout(accessToken string) error {
	return nil
}
