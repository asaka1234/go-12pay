package go_12pay

import "github.com/go-resty/resty/v2"

type Client struct {
	AuthKey     string
	PartnerCode string
	Device      string
	Channel     string

	BaseURL   string
	PayoutURL string

	ryClient *resty.Client
}

func NewClient(authKey string, partnerCode string, device string, channel string, baseURL string, payoutURL string) *Client {
	return &Client{
		AuthKey:     authKey,
		PartnerCode: partnerCode,
		Device:      device,
		Channel:     channel,
		BaseURL:     baseURL,
		PayoutURL:   payoutURL,
		ryClient:    resty.New(), //client实例
	}
}
