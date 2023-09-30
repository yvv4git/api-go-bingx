package bingx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/yvv4git/api-go-bingx/utils"
)

// SwapBalanceRequest - specific implementation of common contract.
type SwapBalanceRequest struct {
	apiURL        string
	apiPath       string
	apiKey        string
	apiSecret     string
	clientTimeout time.Duration
	userAgent     string
}

// NewSwapBalanceRequest - used for create instance of SwapBalanceRequest.
func NewSwapBalanceRequest(apiURL, apiKey, apiSecret string) *SwapBalanceRequest {
	const (
		defaultClientTimeout = time.Second * 2
	)

	return &SwapBalanceRequest{
		apiURL:        apiURL,
		apiPath:       "/openApi/swap/v2/user/balance",
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		clientTimeout: defaultClientTimeout,
		userAgent:     "Mozilla/5.0",
	}
}

// Process - used for create.
func (s *SwapBalanceRequest) Process(ctx context.Context) (*SwapBalanceResponse, error) {
	payloadURL := NewSwapBalancePayload(utils.CurrentTimestamp(), s.apiSecret)
	payloadStr := payloadURL.Create()
	urlPath := fmt.Sprintf("%s%s?%s", s.apiURL, s.apiPath, payloadStr)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error on create request: %w", err)
	}

	request.Header = http.Header{
		"X-Bx-Apikey": []string{s.apiKey},
		"User-Agent":  []string{s.userAgent},
	}

	client := http.Client{
		Timeout: s.clientTimeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, response.Body); err != nil {
		return nil, fmt.Errorf("error on copy body to buffer: %w", err)
	}

	var result SwapBalanceResponse
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error on unmarshal response: %w", err)
	}

	return &result, nil
}

// SetUserAgent - used for setup user agent of http request.
func (s *SwapBalanceRequest) SetUserAgent(userAgent string) *SwapBalanceRequest {
	s.userAgent = userAgent

	return s
}

// SwapBalanceResponse - use as response from api.
//
//nolint:tagliatelle
type SwapBalanceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Balance struct {
			UserID           string `json:"userId"`
			Asset            string `json:"asset"`
			Balance          string `json:"balance"`
			Equity           string `json:"equity"`
			UnrealizedProfit string `json:"unrealizedProfit"`
			RealisedProfit   string `json:"realisedProfit"`
			AvailableMargin  string `json:"availableMargin"`
			UsedMargin       string `json:"usedMargin"`
			FreezedMargin    string `json:"freezedMargin"`
		} `json:"balance"`
	} `json:"data"`
}
