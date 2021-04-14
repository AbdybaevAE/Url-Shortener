package token

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abdybaevae/url-shortener/pkg/conf"
	typederr "github.com/abdybaevae/url-shortener/pkg/errors/typed"
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type service struct {
	conf             *conf.Config
	accessExpiredAt  time.Duration
	refreshExpiredAt time.Duration
}
type TokenDetails struct {
	Access           string
	Refresh          string
	AccessExpiredAt  int64
	RefreshExpiredAt int64
	AccessUuid       string
	RefreshUuid      string
}
type TokenInfo struct {
	UserId int
}

func New(conf *conf.Config) (TokenService, error) {
	accessExpiredAt, err := time.ParseDuration(conf.AccessTokenExpiredAt)
	if err != nil {
		return nil, err
	}
	refreshExpiredAt, err := time.ParseDuration(conf.RefreshTokenExpiredAt)
	if err != nil {
		return nil, err
	}

	return &service{
		accessExpiredAt:  accessExpiredAt,
		refreshExpiredAt: refreshExpiredAt,
		conf:             conf,
	}, nil
}
func (s *service) GenerateToken(user *models.User) (*TokenDetails, error) {
	token := &TokenDetails{
		AccessExpiredAt:  time.Now().Add(s.accessExpiredAt).Unix(),
		RefreshExpiredAt: time.Now().Add(s.refreshExpiredAt).Unix(),
		AccessUuid:       uuid.NewV4().String(),
		RefreshUuid:      uuid.NewV4().String(),
	}
	accessTokenClaims := jwt.MapClaims{
		"exp":    token.AccessExpiredAt,
		"userId": user.Id,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	var err error
	token.Access, err = accessToken.SignedString([]byte(s.conf.AccessTokenSecret))
	if err != nil {
		return nil, err
	}
	refreshTokenClaims := jwt.MapClaims{
		"exp":    token.RefreshExpiredAt,
		"userId": user.Id,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	token.Refresh, err = refreshToken.SignedString([]byte(s.conf.RefreshTokenSecret))
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (s *service) RefreshToken(refreshToken string) (*TokenDetails, error) {
	return nil, nil
}
func (s *service) DeleteToken(accessToken string) error {
	return nil
}
func (s *service) VerifyAccessToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, typederr.InvalidToken
		}
		return []byte(s.conf.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, typederr.InvalidToken
	}
	claims, ok := token.Claims.(jwt.Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	tokenInfo := &TokenInfo{}
	// tokenInfo.UserId, err = strconv.Atoi(claims["user_id"])
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}
