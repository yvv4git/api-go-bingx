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
	/*
		Show my assets.
		Asserts are a set of characters and their number.
	*/
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

	spotAssetsRequest := api.NewSpotAssetsRequest(apiURL, apiKey, apiSecret)
	response, err := spotAssetsRequest.Process(context.Background())
	if err != nil {
		log.Fatalf("error on process request: %v", err)
	}

	spotAssetsResponse, ok := response.(*api.SpotAssetsResponse)
	if !ok {
		log.Fatal("error on convert response to SpotAssetsResponse")
	}

	for _, balance := range spotAssetsResponse.Data.Balances {
		log.Printf("[%s] %v", balance.Asset, balance.Free)
	}
}
