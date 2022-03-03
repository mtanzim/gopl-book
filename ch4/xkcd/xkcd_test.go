package xkcd

import (
	"reflect"
	"testing"
)

func Test_makeURL(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "basic test case",
			args: args{id: "44"},
			want: "https//xkcd.com/44/info.0.json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeURL(tt.args.id); got != tt.want {
				t.Errorf("makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetComic(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "sanity check",
			args:    args{id: "781"},
			want:    781,
			wantErr: false,
		},
		{
			name:    "sanity check 2",
			args:    args{id: "500"},
			want:    500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetComicFromRemote(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Num, tt.want) {
				t.Errorf("GetComic() = %v, want %v", got.Num, tt.want)
			}
		})
	}
}
