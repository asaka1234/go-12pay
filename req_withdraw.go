package go_12pay

import (
	"crypto/tls"
)

// 退款
func (cli *Client) Withdraw(req One2PayWithdrawRequest) (*One2PayWithdrawResponse, error) {

	rawURL := cli.WithdrawURL

	//返回值会放到这里
	var result One2PayWithdrawResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(cli.AuthKey, cli.Channel, cli.Device, cli.PartnerCode)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}
