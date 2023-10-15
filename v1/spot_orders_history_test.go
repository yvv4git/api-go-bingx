package v1_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/yvv4git/api-go-bingx/v1"
	mock "gopkg.in/h2non/gock.v1"
)

func TestSpotOrdersHistoryRequest_Process(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	mock.New(apiURL).Get("/openApi/spot/v1/trade/historyOrders").
		Reply(200).
		JSON(`{
   "code":0,
   "data": {
		"Orders":null
	}
}`)
	defer mock.Off()

	spotOrdersHistoryRequest := api.NewSpotOrdersListRequest(api.SpotOrdersListParams{
		ApiURL:    apiURL,
		ApiKey:    apiKey(),
		ApiSecret: apiSecret(),
	})

	resp, err := spotOrdersHistoryRequest.Process(context.Background())
	require.NoError(t, err)
	require.Empty(t, resp.Data.Orders)
}
