package ad_listing

import (
	"context"
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

func NewClient(baseUrl string, retryTimes int, log *log.Logger) *client {
	// TODO #4 refactor NewClient using functional options
	return &client{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		retryTimes: retryTimes,
		logger:     log,
	}
}

type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
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
	fmt.Printf("\nResponse %v", string(b))

	var adResp AdsResponse
	// TODO #2 unmarshal json

	return &adResp, nil
}

type AdsResponse struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AdId int `json:"ad_id"`
	//TODO #1 Define struct
	// list_id , account_name, subject, list_time
}
