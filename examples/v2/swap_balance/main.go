package main

import (
	"context"
	"log"
	"os"

	api "github.com/yvv4git/api-go-bingx/v2"
)

const (
	envURL    = "API_URL"
	envKey    = "API_KEY"
	envSecret = "API_SECRET"
)

func main() {
	/*
		Show my balance.
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

	swapBalanceRequest := api.NewSwapBalanceRequest(apiURL, apiKey, apiSecret)
	resp, err := swapBalanceRequest.Process(context.Background())
	if err != nil {
		log.Fatalf("error on process request: %v", err)
	}

	log.Printf("%#v", resp)
}
