package keys

import (
	"testing"

	mock_key_repo "github.com/abdybaevae/url-shortener/mocks/pkg/repos/key"
	mock_algo_srv "github.com/abdybaevae/url-shortener/mocks/pkg/services/algo"
	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	type deps struct {
		algoSrv *mock_algo_srv.MockAlgoService
		keyRepo *mock_key_repo.MockKeyRepo
	}
	tt := []struct {
		name    string
		want    string
		wantErr error
	}{
		{name: "Generate key", want: "asjdkjsa"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := &deps{
				algoSrv: mock_algo_srv.NewMockAlgoService(ctrl),
				keyRepo: mock_key_repo.NewMockKeyRepo(ctrl),
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
