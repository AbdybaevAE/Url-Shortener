package server

import (
	"context"

	link_service "github.com/abdybaevae/url-shortener/pkg/services/link"
	pbLink "github.com/abdybaevae/url-shortener/proto/links"
	pbUsers "github.com/abdybaevae/url-shortener/proto/users"
)

// Backend implements the protobuf interface
type Backend struct {
	linkService link_service.LinkService
}

func NewBackend(linkService link_service.LinkService) *Backend {
	return &Backend{
		linkService: linkService,
	}
}
func (b *Backend) ShortenLink(ctx context.Context, in *pbLink.ShortenLinkReq) (*pbLink.ShortenLinkRes, error) {
	shortLink, err := b.linkService.ShortenLink(in.Link)
	if err != nil {
		return nil, err
	}
	return &pbLink.ShortenLinkRes{ShortLink: shortLink}, nil
}
func (b *Backend) GetLink(ctx context.Context, in *pbLink.GetLinkReq) (*pbLink.GetLinkRes, error) {
	link, err := b.linkService.GetLink(in.Key)
	if err != nil {
		return nil, err
	}
	return &pbLink.GetLinkRes{Link: link}, nil
}

func (b *Backend) Register(ctx context.Context, in *pbUsers.RegisterReq) (*pbUsers.RegisterRes, error) {

	return nil, nil
}
func (b *Backend) Login(ctx context.Context, in *pbUsers.LoginReq) (*pbUsers.LoginRes, error) {
	return nil, nil
}
func (b *Backend) RefreshToken(ctx context.Context, in *pbUsers.RefreshTokenReq) (*pbUsers.RefreshTokenRes, error) {
	return nil, nil
}
