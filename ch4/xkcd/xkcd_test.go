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
		{name: "url string sub works",
			args: args{id: "44"},
			want: "https://xkcd.com/44/info.0.json"},
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
		want    string
		wantErr bool
	}{
		{
			name:    "from remote",
			args:    args{id: "781"},
			want:    "remote",
			wantErr: false,
		},
		{
			name:    "from remote again",
			args:    args{id: "782"},
			want:    "remote",
			wantErr: false,
		},
		{
			name:    "now from cache",
			args:    args{id: "781"},
			want:    "cache",
			wantErr: false,
		},
	}
	c := NewCache()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, err := GetComic(tt.args.id, c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getComicFromRemote(t *testing.T) {
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
			name:    "from remote id:781",
			args:    args{id: "781"},
			want:    781,
			wantErr: false,
		},
		{
			name:    "from remote id:500",
			args:    args{id: "500"},
			want:    500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getComicFromRemote(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getComicFromRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Num, tt.want) {
				t.Errorf("getComicFromRemote() = %v, want %v", got.Num, tt.want)
			}
		})
	}
}
