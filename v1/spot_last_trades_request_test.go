package v1_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	api "github.com/yvv4git/api-go-bingx/v1"
	mock "gopkg.in/h2non/gock.v1"
)

func TestSpotLastTradesRequest_preparePayload(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	type args struct {
		timestamp int64
		symbol    string
		limit     int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				timestamp: 1675929864070,
				symbol:    "BTC-USDT",
				limit:     10,
			},
			want: "symbol=BTC-USDT&limit=10&timestamp=1675929864070&" +
				"signature=9c53bf7b2cb28b5ce1565fcc3541d9fcca768c400f40d74d039b99e364024711",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := api.NewSpotLastTradesRequest(apiURL, apiKey(), apiSecret())
			result := instance.PreparePayload(tt.args.timestamp, tt.args.symbol, tt.args.limit)
			t.Logf("--->%v", result)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestSpotLastTradesRequest_Process(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	mock.New(apiURL).Get("/openApi/spot/v1/market/trades").
		Reply(200).
		JSON(`{
   "code":0,
   "data":[
      {
         "id":53218856,
         "price":27907.74,
         "qty":0.07238,
         "time":1680980556877,
         "buyerMaker":false
      },
      {
         "id":53218855,
         "price":27903.61,
         "qty":0.07772,
         "time":1680980553940,
         "buyerMaker":true
      },
      {
         "id":53218854,
         "price":27905.64,
         "qty":0.07702,
         "time":1680980551209,
         "buyerMaker":false
      },
      {
         "id":53218853,
         "price":27902.64,
         "qty":0.05545,
         "time":1680980546604,
         "buyerMaker":true
      },
      {
         "id":53218852,
         "price":27879.64,
         "qty":0.06465,
         "time":1680980538352,
         "buyerMaker":true
      },
      {
         "id":53218851,
         "price":27878.62,
         "qty":0.07002,
         "time":1680980536407,
         "buyerMaker":false
      },
      {
         "id":53218850,
         "price":27875.73,
         "qty":0.06074,
         "time":1680980531475,
         "buyerMaker":true
      },
      {
         "id":53218849,
         "price":27878.10,
         "qty":0.07894,
         "time":1680980527776,
         "buyerMaker":false
      },
      {
         "id":53218848,
         "price":27880.94,
         "qty":0.05798,
         "time":1680980525391,
         "buyerMaker":true
      },
      {
         "id":53218847,
         "price":27877.85,
         "qty":0.06178,
         "time":1680980522518,
         "buyerMaker":false
      }
   ]
}`)
	defer mock.Off()

	spotLastTradesRequest := api.NewSpotLastTradesRequest(apiURL, apiKey(), apiSecret())
	spotLastTradesRequest.SetUserAgent("Mozilla/6.0")
	spotLastTradesRequest.SetSymbol("BTC-USDT")
	spotLastTradesRequest.SetLimit(10)

	resp, err := spotLastTradesRequest.Process(context.Background())
	require.NoError(t, err)

	require.Equal(t, 10, len(resp.Data))
}
