package bingx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/yvv4git/api-go-bingx"
)

func TestSpotDepthPayload_Create(t *testing.T) {
	type fields struct {
		timestamp int64
		limit     int
		symbol    string
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
				limit:     9,
				symbol:    "OLE-USDT",
				signature: "BLVucNV235GppMNF777MbqQMueKw8LX58EUNKGH32ZWtOOa1msO5IpChY8Q8hLyLXjilLXIKC2Xf3St01dV",
			},
			want: "symbol=OLE-USDT&" +
				"limit=9&" +
				"timestamp=1675929864070&signature=0c05cf8989bdd1f4573d4f301e7e93a567f43ba476d4f4dc087ffbd8d1c16505",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := api.NewSpotDepthPayload(tt.fields.timestamp, tt.fields.limit, tt.fields.symbol, tt.fields.signature)
			result := instance.Create()
			assert.Equal(t, tt.want, result)
		})
	}
}
