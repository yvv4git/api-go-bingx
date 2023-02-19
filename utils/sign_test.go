package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bingx "github.com/yvv4git/api-go-bingx/utils"
)

func Test_sign(t *testing.T) {
	type args struct {
		secrecy string
		payload string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				secrecy: "BLVucNV235GppMNF777MbqQMueKw8LX58EUNKGH32ZWtOOa1msO5IpChY8Q8hLyLXjilLXIKC2Xf3St01dV",
				payload: "1675931228479",
			},
			want: "97511c171762b1ad8eeef01611ee73ac6a9377b7ffe55bee8de51e61f7a4d0e4",
		},
		{
			name: "CASE-2",
			args: args{
				secrecy: "BLVucNV235GppMNF777MbqQMueKw8LX58EUNKGH32ZWtOOa1msO5IpChY8Q8hLyLXjilLXIKC2Xf3St01dV",
				payload: "timestamp=1676199978197",
			},
			want: "886f63d49c8576d7c821b7609ae8a13d7655defaf4d8b1945ca88a76527dbc70",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bingx.Sign(tt.args.secrecy, tt.args.payload)
			assert.Equal(t, tt.want, result)
		})
	}
}
