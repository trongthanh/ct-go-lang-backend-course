package ad_listing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)

type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}

type option func(*client)

func NewClient(opt ...option) *client {
	// TODO #4 refactor NewClient using functional options ✅
	c := new(client)
	for _, o := range opt {
		o(c)
	}

	// Set a default client if one was not provided
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}

func WithBaseUrl(url string) option {
	return func(c *client) {
		c.baseUrl = url
	}
}

func WithLogger(logger *log.Logger) option {
	return func(c *client) {
		c.logger = logger
	}
}

func WithRetryTimes(times int) option {
	return func(c *client) {
		c.retryTimes = times
	}
}

func (c *client) GetAdByCate(ctx context.Context, cate string) (*AdsResponse, error) {
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request - Cate %v, Duration: %v", cate, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, cate)

	// TODO #3 implement retry if StatusCode = 5xx
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("\nResponse %v", string(b))

	var adResp AdsResponse
	// TODO #2 unmarshal json ✅
	unmarshalErr := json.Unmarshal(b, &adResp)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &adResp, nil
}

type AdsResponse struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AdId int `json:"ad_id"`
	// TODO #1 define struct ✅
	ListId      int    `json:"list_id"`
	AccountName string `json:"account_name"`
	Subject     string `json:"subject"`
	ListTime    int64  `json:"list_time"`
}
