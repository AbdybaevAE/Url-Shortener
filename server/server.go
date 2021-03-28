package server

import (
	"context"

	"github.com/abdybaevae/url-shortener/pkg/services/links"
	pbLink "github.com/abdybaevae/url-shortener/proto"
	// ""
)

// Backend implements the protobuf interface
type Backend struct {
	linkService links.LinkService
}

func NewBackend(linkService links.LinkService) *Backend {
	return &Backend{
		linkService: linkService,
	}
}
func (b *Backend) Shorten(ctx context.Context, in *pbLink.ShortenReq) (*pbLink.ShortenRes, error) {
	shortLink, err := b.linkService.Shorten(in.Link)
	if err != nil {
		return nil, err
	}
	return &pbLink.ShortenRes{ShortLink: shortLink}, nil
}
func (b *Backend) GetOriginalFromShorten(ctx context.Context, in *pbLink.GetOriginalFromShortenReq) (*pbLink.GetOriginalFromShortenRes, error) {
	originalLink, err := b.linkService.GetOriginalFromShorten(in.ShortLink)
	if err != nil {
		return nil, err
	}
	return &pbLink.GetOriginalFromShortenRes{OriginalLink: originalLink}, nil
}
