package bingx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/yvv4git/api-go-bingx"
)

func TestSwapBalancePayloadURL_Create(t *testing.T) {
	type fields struct {
		timestamp int64
		signature string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "CASE-1",
			fields: fields{
				timestamp: 1675929864070,
				signature: "BLVucNV235GppMNF777MbqQMueKw8LX58EUNKGH32ZWtOOa1msO5IpChY8Q8hLyLXjilLXIKC2Xf3St01dV",
			},
			want: "timestamp=1675929864070&signature=25680507ab870a5b793b582b8ef10f73dc432afc5bb6e2b60af2aebc073e3f23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := api.NewSwapBalancePayload(tt.fields.timestamp, tt.fields.signature)
			result := instance.Create()
			assert.Equal(t, tt.want, result)
		})
	}
}
