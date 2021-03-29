package keys

import (
	"testing"

	mock_key_repo "github.com/abdybaevae/url-shortener/mocks/pkg/repos/key"
	mock_algo_srv "github.com/abdybaevae/url-shortener/mocks/pkg/services/algo"
	"github.com/abdybaevae/url-shortener/pkg/models"
	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	type deps struct {
		algoSrv *mock_algo_srv.MockAlgoService
		keyRepo *mock_key_repo.MockKeyRepo
		// keySrv
	}
	tt := []struct {
		name    string
		want    string
		wantErr error
		prepare func(f *deps)
	}{
		{
			name: "Generate key",
			want: "some_key",
			prepare: func(f *deps) {
				f.algoSrv.EXPECT().Entity().Return(&models.Algo{Id: 1})
				f.keyRepo.EXPECT().DeleteOne(1).Return("some_key", nil)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := &deps{
				algoSrv: mock_algo_srv.NewMockAlgoService(ctrl),
				keyRepo: mock_key_repo.NewMockKeyRepo(ctrl),
			}
			if tc.prepare != nil {
				tc.prepare(d)
			}
			keyService := New(d.keyRepo, d.algoSrv)
			got, gotErr := keyService.Get()
			if tc.wantErr != nil {
				if tc.wantErr != gotErr {
					t.Errorf("want %v got %v", tc.wantErr, gotErr)
				}
				return
			}
			if got != tc.want {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}
