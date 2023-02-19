package bingx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/yvv4git/api-go-bingx"
	mock "gopkg.in/h2non/gock.v1"
)

func TestSwapBalanceRequest_Create(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	mock.New(apiURL).Get("/openApi/swap/v2/user/balance").
		Reply(200).
		JSON(`{
	   "code":0,
	   "msg":"",
	   "data":{
		  "balance":{
			 "userId":"1327340714010876123",
			 "asset":"USDT",
			 "balance":"0.0000",
			 "equity":"0.0000",
			 "unrealizedProfit":"0.0000",
			 "realisedProfit":"0.0000",
			 "availableMargin":"0.0000",
			 "usedMargin":"0.0000",
			 "freezedMargin":"0.0000"
		  }
	   }
	}`)
	defer mock.Off()

	swapBalanceRequest := api.NewSwapBalanceRequest(apiURL, apiKey(), apiSecret())
	resp, err := swapBalanceRequest.Process(context.Background())
	require.NoError(t, err)

	swapBalanceResp, ok := resp.(*api.SwapBalanceResponse)
	require.True(t, ok)
	require.Equal(t, 0, swapBalanceResp.Code)
	require.Equal(t, "USDT", swapBalanceResp.Data.Balance.Asset)
}

func apiKey() string {
	return "eeDeet5auchohngi8choo6uth3ewingu5gaehaepoh3ahtohqu9cev1jameRoh4aix1beiT1Eiwujogh4"
}

func apiSecret() string {
	return "Iexo0kaek9EWainah4phee3zei6Zej5woow7queixee1cioreethohquo5uW5jei7Aim6doh6che6ooqu7p"
}
