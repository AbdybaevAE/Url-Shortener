package server

import (
	"context"

	link_service "github.com/abdybaevae/url-shortener/pkg/services/link"
	pbLink "github.com/abdybaevae/url-shortener/proto"
)

// Backend implements the protobuf interface
type Backend struct {
	linkService link_service.LinkService
}

func NewBackend(linkService link_service.LinkService) pbLink.LinkServiceServer {
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

func (b *Backend) VisitByKey(ctx context.Context, in *pbLink.VisitByKeyReq) (*pbLink.VisitByKeyRes, error) {
	if err := b.linkService.VisitByKey(in.Key); err != nil {
		return nil, err
	}
	return &pbLink.VisitByKeyRes{}, nil
}
