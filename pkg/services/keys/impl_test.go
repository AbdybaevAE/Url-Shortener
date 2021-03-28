package keys

import "testing"

func TestGet(t *testing.T) {
	tt := []struct {
		name    string
		want    string
		wantErr error
	}{
		{name: "Generate key", want: "asjdkjsa"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			keyService := NewService()
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
