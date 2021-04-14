package user

import (
	http_err "github.com/abdybaevae/url-shortener/pkg/errors/http"
	"github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	usrrepo "github.com/abdybaevae/url-shortener/pkg/repos/user"
	cachesrv "github.com/abdybaevae/url-shortener/pkg/services/cache"
	tokensrv "github.com/abdybaevae/url-shortener/pkg/services/token"
)

type service struct {
	repo     usrrepo.UserRepo
	tokenSrv tokensrv.TokenService
	cacheSrv cachesrv.CacheService
}

func New(
	repo usrrepo.UserRepo,
	tokenSrv tokensrv.TokenService,
	cacheSrv cachesrv.CacheService) UserService {
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
func (s *service) Login(account, password string) (*tokensrv.TokenPair, error) {
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
	tokenDetails, err := s.tokenSrv.GenerateToken(user)
	if err := s.cacheSrv.SetToken(user, tokenDetails); err != nil {
		return nil, err
	}
	return &tokensrv.TokenPair{
		Access:  tokenDetails.Access,
		Refresh: tokenDetails.Refresh,
	}, nil
}
func (s *service) RefreshToken(refreshToken string) (p *tokensrv.TokenPair, err error) {

	return nil, nil
}
func (s *service) Logout(accessToken string) error {
	s.cacheSrv.DeleteTokenUuid()
	return nil
}
