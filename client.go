package go_12pay

import (
	"github.com/asaka1234/go-12pay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params One2PayInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params One2PayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
