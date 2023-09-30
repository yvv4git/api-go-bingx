package bingx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/yvv4git/api-go-bingx"
	mock "gopkg.in/h2non/gock.v1"
)

func TestSpotAssetsRequest_Create(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	mock.New(apiURL).Get("/openApi/spot/v1/account/balance").
		Reply(200).
		JSON(`{
		  "code": 0,
		  "msg": "",
		  "debugMsg": "",
		  "data": {
			"balances": [
			  {
				"asset": "BTC",
				"free": "0.0000091",
				"locked": "0"
			  },
			  {
				"asset": "USDT",
				"free": "4.205595722882512",
				"locked": "0"
			  },
			  {
				"asset": "DOT",
				"free": "0.000732",
				"locked": "0"
			  },
			  {
				"asset": "SOL",
				"free": "0.0000331",
				"locked": "0"
			  },
			  {
				"asset": "NEAR",
				"free": "0.502927",
				"locked": "0"
			  },
			  {
				"asset": "1INCH",
				"free": "0.00594",
				"locked": "0"
			  },
			  {
				"asset": "OLE",
				"free": "0.012600000000020373",
				"locked": "1728.8"
			  },
			  {
				"asset": "APT",
				"free": "0.0000892",
				"locked": "0"
			  }
			]
		  }
		}`)
	defer mock.Off()

	spotAssetsRequest := api.NewSpotAssetsRequest(apiURL, apiKey(), apiSecret())
	resp, err := spotAssetsRequest.Process(context.Background())
	require.NoError(t, err)
	require.Equal(t, "BTC", resp.Data.Balances[0].Asset)
	require.Equal(t, "0.0000091", resp.Data.Balances[0].Free)
}
