package server

import (
	"context"
	"testing"

	mock_link_srv "github.com/abdybaevae/url-shortener/mocks/pkg/services/link"
	"github.com/abdybaevae/url-shortener/mocks_data"
	"github.com/abdybaevae/url-shortener/pkg/errors/http"
	pbLink "github.com/abdybaevae/url-shortener/proto"
	"github.com/golang/mock/gomock"
)

func TestShorten(t *testing.T) {
	type deps struct {
		linkSrv *mock_link_srv.MockLinkService
	}
	tt := []struct {
		name    string
		args    *pbLink.ShortenReq
		want    *pbLink.ShortenRes
		wantErr error
		prepare func(d *deps)
	}{
		{
			name: "Shorten link",
			args: mocks_data.BasicShoretenLinkReq,
			want: mocks_data.BasicShoretenLinkRes,
			prepare: func(d *deps) {
				d.linkSrv.EXPECT().Shorten(mocks_data.BasicLink).Return(mocks_data.BasicLinkShortenedKey, nil)
			},
		},
		{
			name:    "InvalidLink link",
			args:    mocks_data.InvalidShoretenLinkReq,
			wantErr: http.InvalidLink,
			prepare: func(d *deps) {
				d.linkSrv.EXPECT().Shorten(mocks_data.InvalidLink).Return("", http.InvalidLink)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := &deps{
				linkSrv: mock_link_srv.NewMockLinkService(ctrl),
			}
			if tc.prepare != nil {
				tc.prepare(d)
			}
			be := NewBackend(d.linkSrv)
			got, gotErr := be.Shorten(context.Background(), tc.args)

			if tc.wantErr != nil {
				if tc.wantErr != gotErr {
					t.Errorf("want %v got %v", tc.wantErr, gotErr)
				}
				return

			}
			if got.ShortLink != tc.want.ShortLink {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})

	}
}
