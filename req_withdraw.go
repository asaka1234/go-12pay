package go_12pay

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
)

// 退款
// 这个没有回调. 所以主要是看状态
func (cli *Client) Withdraw(req One2PayWithdrawRequest) (*One2PayWithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	//返回值会放到这里
	var result One2PayWithdrawResponse

	textBody, _ := json.Marshal(req)

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(string(textBody)).
		SetHeaders(getPayoutAuthHeaders(cli.Params.PartnerCode, cli.Params.AuthKey, cli.Params.Channel, cli.Params.Device)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetLogger(cli.logger).
		SetError(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d, err:%s", resp.StatusCode(), result.Error)
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	//-----------错误处理------------------------
	if result.Status != 1000 {
		//说明失败
		return nil, errors.New(result.Message)
	}

	//没有error就证明成功了
	return &result, nil
}
