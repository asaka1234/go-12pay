package go_12pay

import "github.com/go-resty/resty/v2"

type Client struct {
	Token     string
	BaseURL   string
	PayoutURL string
	ryClient  *resty.Client
}

func NewClient(token string, baseURL string, payoutURL string) *Client {
	return &Client{
		Token:     token,
		BaseURL:   baseURL,
		PayoutURL: payoutURL,
		ryClient:  resty.New(), //client实例
	}
}
