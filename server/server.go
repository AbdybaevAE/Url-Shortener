package server

import (
	"context"

	"github.com/abdybaevae/url-shortener/pkg/lib/resp"
	link_srv "github.com/abdybaevae/url-shortener/pkg/services/link"
	usr_srv "github.com/abdybaevae/url-shortener/pkg/services/user"
	pbLink "github.com/abdybaevae/url-shortener/proto/links"
	pbUsers "github.com/abdybaevae/url-shortener/proto/users"
)

// Backend implements the protobuf interface
type Backend struct {
	linkSrv link_srv.LinkService
	userSrv usr_srv.UserService
}

func NewBackend(linkSrv link_srv.LinkService, userSrv usr_srv.UserService) *Backend {
	return &Backend{
		linkSrv: linkSrv,
		userSrv: userSrv,
	}
}
func (b *Backend) ShortenLink(ctx context.Context, in *pbLink.ShortenLinkReq) (*pbLink.ShortenLinkRes, error) {
	shortLink, err := b.linkSrv.ShortenLink(in.Link)
	if err != nil {
		return nil, err
	}
	return &pbLink.ShortenLinkRes{ShortLink: shortLink}, nil
}
func (b *Backend) GetLink(ctx context.Context, in *pbLink.GetLinkReq) (*pbLink.GetLinkRes, error) {
	link, err := b.linkSrv.GetLink(in.Key)
	if err != nil {
		return nil, err
	}
	return &pbLink.GetLinkRes{Link: link}, nil
}

func (b *Backend) Register(ctx context.Context, in *pbUsers.RegisterReq) (*pbUsers.RegisterRes, error) {
	if err := b.userSrv.Register(in.Account, in.Password); err != nil {
		return nil, err
	}
	return &pbUsers.RegisterRes{
		Status: resp.Success.String(),
	}, nil
}
func (b *Backend) Login(ctx context.Context, in *pbUsers.LoginReq) (*pbUsers.LoginRes, error) {
	pair, err := b.userSrv.Login(in.Account, in.Password)
	if err != nil {
		return nil, err
	}
	return &pbUsers.LoginRes{
		AccessToken:  pair.Access,
		RefreshToken: pair.Refresh,
	}, nil
}
func (b *Backend) RefreshToken(ctx context.Context, in *pbUsers.RefreshTokenReq) (*pbUsers.RefreshTokenRes, error) {
	pair, err := b.userSrv.RefreshToken(in.ResreshToken)
	if err != nil {
		return nil, err
	}
	return &pbUsers.RefreshTokenRes{
		AccessToken:  pair.Access,
		RefreshToken: pair.Refresh,
	}, nil
}
