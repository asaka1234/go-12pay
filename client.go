package go_12pay

import (
	"github.com/asaka1234/go-12pay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	PartnerCode string //商户号
	AuthKey     string //accessKey
	Device      string
	Channel     string

	BaseURL   string //充值url
	PayoutURL string //退款url

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, partnerCode string, authKey string, device string, channel string, baseURL string, payoutURL string) *Client {
	return &Client{
		AuthKey:     authKey,
		PartnerCode: partnerCode,
		Device:      device,
		Channel:     channel,
		BaseURL:     baseURL,
		PayoutURL:   payoutURL,
		ryClient:    resty.New(), //client实例
		logger:      logger,
	}
}
