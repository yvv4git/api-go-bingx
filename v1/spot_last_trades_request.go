package v1

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

// SpotLastTradesRequest - specific implementation of common contract.
type SpotLastTradesRequest struct {
	apiURL        string
	apiPath       string
	apiKey        string
	apiSecret     string
	clientTimeout time.Duration
	userAgent     string
	symbol        string
	limit         int
}

// SpotLastTradesResponse - use as response from api.
//
//nolint:tagliatelle
type SpotLastTradesResponse struct {
	Code int `json:"code"`
	Data []struct {
		ID         int     `json:"id"`
		Price      float64 `json:"price"`
		Qty        float64 `json:"qty"`
		Time       int64   `json:"time"`
		BuyerMaker bool    `json:"buyerMaker"`
	} `json:"data"`
}

// NewSpotLastTradesRequest - used for create instance of SwapBalanceRequest.
func NewSpotLastTradesRequest(apiURL, apiKey, apiSecret string) *SpotLastTradesRequest {
	const (
		defaultPath          = "/openApi/spot/v1/market/trades"
		defaultClientHTTP    = "Mozilla/5.0"
		defaultClientTimeout = time.Second * 2
		defaultSymbol        = "BTC-USDT"
		defaultLimit         = 9
	)

	return &SpotLastTradesRequest{
		apiURL:        apiURL,
		apiPath:       defaultPath,
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		clientTimeout: defaultClientTimeout,
		userAgent:     defaultClientHTTP,
		symbol:        defaultSymbol,
		limit:         defaultLimit,
	}
}

func (s *SpotLastTradesRequest) PreparePayload(timestamp int64, symbol string, limit int) string {
	payload := fmt.Sprintf("symbol=%s&limit=%d&timestamp=%d", symbol, limit, timestamp)
	signature := utils.Sign(s.apiSecret, payload)

	return fmt.Sprintf("%s&signature=%s", payload, signature)
}

// SetUserAgent - used for setup user agent of http request.
func (s *SpotLastTradesRequest) SetUserAgent(userAgent string) {
	if userAgent != "" {
		s.userAgent = userAgent
	}
}

func (s *SpotLastTradesRequest) SetSymbol(symbol string) {
	s.symbol = symbol
}

// SetLimit - used for setup non-default limit.
func (s *SpotLastTradesRequest) SetLimit(limit int) {
	s.limit = limit
}

// Process - used for processing request.
func (s *SpotLastTradesRequest) Process(ctx context.Context) (*SpotLastTradesResponse, error) {
	payloadStr := s.PreparePayload(utils.CurrentTimestamp(), s.symbol, s.limit)
	urlPath := fmt.Sprintf("%s%s?%s", s.apiURL, s.apiPath, payloadStr)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error on create request: %w", err)
	}

	request.Header = http.Header{
		"X-Bx-Apikey":  []string{s.apiKey},
		"User-Agent":   []string{s.userAgent},
		"Content-Type": []string{"application/json"},
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

	var result SpotLastTradesResponse
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error on unmarshal response: %w", err)
	}

	return &result, nil
}
