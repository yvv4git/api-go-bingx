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

// SpotDepthRequest - specific implementation of common contract.
type SpotDepthRequest struct {
	apiURL        string
	apiPath       string
	apiKey        string
	apiSecret     string
	clientTimeout time.Duration
	userAgent     string
	symbol        string
	limit         int
}

// SpotDepthResponse - used as message from api.
type SpotDepthResponse struct {
	Code int `json:"code"`
	Data struct {
		Bids [][]string `json:"bids"`
		Asks [][]string `json:"asks"`
	} `json:"data"`
}

// NewSpotDepthRequest - used for create instance of SpotDepthRequest.
func NewSpotDepthRequest(apiURL, apiKey, apiSecret string) *SpotDepthRequest {
	const (
		defaultPath          = "/openApi/spot/v1/market/depth"
		defaultClientHTTP    = "Mozilla/5.0"
		defaultClientTimeout = time.Second * 2
		defaultSymbol        = "OLE-USDT"
		defaultLimit         = 9
	)

	return &SpotDepthRequest{
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

// Process - used for create.
func (s *SpotDepthRequest) Process(ctx context.Context) (*SpotDepthResponse, error) {
	timestamp := utils.CurrentTimestamp()
	urlData := fmt.Sprintf("symbol=%s&limit=%d&timestamp=%d", s.symbol, s.limit, timestamp)

	payload := fmt.Sprintf(
		"%s&signature=%s",
		urlData,
		utils.Sign(s.apiSecret, urlData))
	urlPath := fmt.Sprintf("%s%s?%s", s.apiURL, s.apiPath, payload)

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

	var result SpotDepthResponse
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error on unmarshal response: %w", err)
	}

	return &result, nil
}

// SetUserAgent - used for setup user agent of http request.
func (s *SpotDepthRequest) SetUserAgent(userAgent string) *SpotDepthRequest {
	s.userAgent = userAgent

	return s
}

// SetSymbol - used for setup non-default symbol.
func (s *SpotDepthRequest) SetSymbol(symbol string) *SpotDepthRequest {
	s.symbol = symbol

	return s
}

// SetLimit - used for setup non-default limit.
func (s *SpotDepthRequest) SetLimit(limit int) *SpotDepthRequest {
	s.limit = limit

	return s
}
