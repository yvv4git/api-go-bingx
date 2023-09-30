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

// SpotAssetsRequest - specific implementation of common contract.
type SpotAssetsRequest struct {
	apiURL        string
	apiPath       string
	apiKey        string
	apiSecret     string
	clientTimeout time.Duration
	userAgent     string
}

// NewSpotAssetsRequest - used for create instance of SpotAssetsRequest.
func NewSpotAssetsRequest(apiURL, apiKey, apiSecret string) *SpotAssetsRequest {
	const (
		defaultPath          = "/openApi/spot/v1/account/balance"
		defaultClientHTTP    = "Mozilla/5.0"
		defaultClientTimeout = time.Second * 2
	)

	return &SpotAssetsRequest{
		apiURL:        apiURL,
		apiPath:       defaultPath,
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		clientTimeout: defaultClientTimeout,
		userAgent:     defaultClientHTTP,
	}
}

// Process - used for create.
func (s *SpotAssetsRequest) Process(ctx context.Context) (*SpotAssetsResponse, error) {
	timestamp := utils.CurrentTimestamp()
	signature := utils.Sign(s.apiSecret, fmt.Sprintf("timestamp=%d", timestamp))
	payloadStr := fmt.Sprintf("timestamp=%d&signature=%s", timestamp, signature)
	urlPath := fmt.Sprintf("%s%s?%s", s.apiURL, s.apiPath, payloadStr)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("error on create request: %w", err)
	}

	request.Header = http.Header{
		"X-Bx-Apikey":  []string{s.apiKey},
		"User-Agent":   []string{s.userAgent},
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"Accept"},
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

	var result SpotAssetsResponse
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error on unmarshal response: %w", err)
	}

	return &result, nil
}

// SetUserAgent - used for setup user agent of http request.
func (s *SpotAssetsRequest) SetUserAgent(userAgent string) *SpotAssetsRequest {
	s.userAgent = userAgent

	return s
}

// SpotAssetsResponse - used as message from api.
//
//nolint:tagliatelle
type SpotAssetsResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debugMsg"`
	Data     struct {
		Balances []struct {
			Asset  string `json:"asset"`
			Free   string `json:"free"`
			Locked string `json:"locked"`
		} `json:"balances"`
	} `json:"data"`
}
