package go_12pay

import (
	"crypto/tls"
)

// 退款
func (cli *Client) Payout(req WithdrawRequest) (*WithdrawResponse, error) {

	reqPath := "/payout"
	rawURL := cli.PayoutURL + reqPath

	//返回值会放到这里
	var result WithdrawResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(HEAD_AUTH_KEY, HEAD_CHANNEL, HEAD_DEVICE, HEAD_PARTNER_CODE)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}
