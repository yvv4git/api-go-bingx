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

type SpotOrdersHistoryRequest struct {
	apiURL        string
	apiPath       string
	apiKey        string
	apiSecret     string
	clientTimeout time.Duration
	userAgent     string
	symbol        string
	limit         int
}

type SpotOrdersHistoryResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debugMsg"`
	Data     struct {
		Orders []struct {
			Symbol              string `json:"symbol"`
			OrderID             int64  `json:"orderId"`
			Price               string `json:"price"`
			OrigQty             string `json:"origQty"`
			ExecutedQty         string `json:"executedQty"`
			CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
			Status              string `json:"status"`
			Type                string `json:"type"`
			Side                string `json:"side"`
			Time                int64  `json:"time"`
			UpdateTime          int64  `json:"updateTime"`
			OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
		} `json:"orders"`
	} `json:"data"`
}

type SpotOrdersListParams struct {
	ApiURL, ApiKey, ApiSecret, Symbol string
	Limit                             int
}

func NewSpotOrdersListRequest(p SpotOrdersListParams) *SpotOrdersHistoryRequest {
	const (
		defaultPath          = "/openApi/spot/v1/trade/historyOrders"
		defaultClientHTTP    = "Mozilla/5.0"
		defaultClientTimeout = time.Second * 2
		defaultSymbol        = "BTC-USDT"
		defaultLimit         = 9
	)

	if p.Symbol == "" {
		p.Symbol = defaultSymbol
	}

	if p.Limit == 0 {
		p.Limit = defaultLimit
	}

	return &SpotOrdersHistoryRequest{
		apiURL:        p.ApiURL,
		apiPath:       defaultPath,
		apiKey:        p.ApiKey,
		apiSecret:     p.ApiSecret,
		clientTimeout: defaultClientTimeout,
		userAgent:     defaultClientHTTP,
		symbol:        p.Symbol,
		limit:         p.Limit,
	}
}

func (s *SpotOrdersHistoryRequest) Process(ctx context.Context) (*SpotOrdersHistoryResponse, error) {
	timeStamp := utils.CurrentTimestamp()
	sigStr := fmt.Sprintf("symbol=%s&timestamp=%d", s.symbol, timeStamp)
	signature := utils.Sign(s.apiSecret, sigStr)
	paramStr := fmt.Sprintf("symbol=%s&timestamp=%d&signature=%s", s.symbol, timeStamp, signature)
	urlPath := fmt.Sprintf("%s%s?%s", s.apiURL, s.apiPath, paramStr)

	request, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	request.Header = http.Header{
		"X-Bx-Apikey":  []string{s.apiKey},
		"User-Agent":   []string{"Mozilla/5.0"},
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

	var result SpotOrdersHistoryResponse
	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error on unmarshal response: %w", err)
	}

	return &result, nil
}
