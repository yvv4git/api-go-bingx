package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/yvv4git/api-go-bingx/v1"
)

const (
	envURL    = "API_URL"
	envKey    = "API_KEY"
	envSecret = "API_SECRET"
)

func main() {
	/*
		Show orders list.
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

	spotLastTradesRequest := api.NewSpotOrdersListRequest(api.SpotOrdersListParams{
		ApiURL:    apiURL,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		Symbol:    "OLE-USDT",
		Limit:     10,
	})

	response, err := spotLastTradesRequest.Process(context.Background())
	if err != nil {
		log.Fatalf("error on process request: %v", err)
	}

	fmt.Printf("Code: %#v \n", response.Code)
	for _, order := range response.Data.Orders {
		fmt.Printf(
			"ID:%v Symbol:%v Price:%v Type:%v Status:%v \n",
			order.OrderID, order.Symbol, order.Price, order.Type, order.Status)
	}
}
