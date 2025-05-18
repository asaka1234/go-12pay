package go_12pay

import (
	"crypto/tls"
)

// 生成支付二维码
func (cli *Client) GenQRCode(req One2PayGenQRCodeRequest) (*One2PayGenQRCodeResponse, error) {

	rawURL := cli.DepositURL

	//返回值会放到这里
	var result One2PayGenQRCodeResponse

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
