package main

import (
	"context"
	"fmt"
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
		Show last trades.
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

	spotLastTradesRequest := api.NewSpotLastTradesRequest(apiURL, apiKey, apiSecret)
	spotLastTradesRequest.SetSymbol("BTC-USDT")
	spotLastTradesRequest.SetLimit(10)

	response, err := spotLastTradesRequest.Process(context.Background())
	if err != nil {
		log.Fatalf("error on process request: %v", err)
	}

	fmt.Printf("Code: %#v \n", response.Code)
	for _, value := range response.Data {
		fmt.Printf(
			"ID=%v Price=%v Qty=%v BuyerMaker=%v TimeStamp=%v \n",
			value.ID, value.Price, value.Qty, value.BuyerMaker, value.Time,
		)
	}
}
