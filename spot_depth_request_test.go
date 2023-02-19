package bingx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/yvv4git/api-go-bingx"
	mock "gopkg.in/h2non/gock.v1"
)

func TestSpotDepthRequest_Create(t *testing.T) {
	const (
		apiURL = "http://localhost:9119"
	)

	mock.New(apiURL).Get("/openApi/spot/v1/market/depth").
		Reply(200).
		JSON(`{
		  "code": 0,
		  "data": {
			"bids": [
			  [
				"0.03899",
				"37.6"
			  ],
			  [
				"0.03895",
				"40.4"
			  ],
			  [
				"0.03892",
				"91.9"
			  ],
			  [
				"0.03891",
				"149.8"
			  ],
			  [
				"0.03890",
				"81.3"
			  ],
			  [
				"0.03886",
				"38.4"
			  ],
			  [
				"0.03885",
				"43.4"
			  ],
			  [
				"0.03881",
				"47.2"
			  ],
			  [
				"0.03878",
				"35.5"
			  ]
			],
			"asks": [
			  [
				"0.03928",
				"36.1"
			  ],
			  [
				"0.03925",
				"45.2"
			  ],
			  [
				"0.03923",
				"37.2"
			  ],
			  [
				"0.03922",
				"43.1"
			  ],
			  [
				"0.03920",
				"44.6"
			  ],
			  [
				"0.03919",
				"47.4"
			  ],
			  [
				"0.03918",
				"41.2"
			  ],
			  [
				"0.03916",
				"85.5"
			  ],
			  [
				"0.03914",
				"34.4"
			  ]
			]
		  }
		}`)
	defer mock.Off()

	spotDepthRequest := api.NewSpotDepthRequest(apiURL, apiKey(), apiSecret())
	spotDepthRequest.SetSymbol("OLE-USDT")
	spotDepthRequest.SetLimit(9)

	resp, err := spotDepthRequest.Process(context.Background())
	require.NoError(t, err)

	spotDepthResp, ok := resp.(*api.SpotDepthResponse)
	require.True(t, ok)
	require.Equal(t, 9, len(spotDepthResp.Data.Asks))
	require.Equal(t, 9, len(spotDepthResp.Data.Bids))
}
