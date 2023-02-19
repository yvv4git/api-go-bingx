package main

import (
	"context"
	"log"
	"os"

	api "github.com/yvv4git/api-go-bingx"
)

const (
	envURL    = "API_URL"
	envKey    = "API_KEY"
	envSecret = "API_SECRET"
)

func main() {
	getEnvOrPanic := func(key string) string {
		value, ok := os.LookupEnv(key)
		if !ok {
			log.Panicf("failed to read environment key: %s", key)
		}

		return value
	}

	apiURL := getEnvOrPanic(envURL)
	apiKey := getEnvOrPanic(envKey)
	apiSecret := getEnvOrPanic(envSecret)

	spotDepthRequest := api.NewSpotDepthRequest(apiURL, apiKey, apiSecret)
	spotDepthRequest.SetSymbol("BTC-USDT")
	spotDepthRequest.SetLimit(10)

	response, err := spotDepthRequest.Process(context.Background())
	if err != nil {
		log.Fatalf("error on process request: %v", err)
	}

	spotDepthResponse, ok := response.(*api.SpotDepthResponse)
	if !ok {
		log.Fatal("error on convert response to SpotAssetsResponse")
	}

	for _, value := range spotDepthResponse.Data.Bids {
		log.Printf("Bid: %v - %v", value[0], value[1])
	}

	log.Println("-----")

	for _, value := range spotDepthResponse.Data.Asks {
		log.Printf("Ask: %v - %v", value[0], value[1])
	}
}
