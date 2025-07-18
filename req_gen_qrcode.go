package go_12pay

import (
	"crypto/tls"
	"errors"
	"github.com/asaka1234/go-12pay/utils"
	jsoniter "github.com/json-iterator/go"
)

// 生成支付二维码
func (cli *Client) GenQRCode(req One2PayGenQRCodeRequest) (*One2PayGenQRCodeResponse, error) {

	rawURL := cli.Params.DepositUrl

	//ref4里放的就是签名字符串
	req.Ref4 = utils.GenSign(req.Amount, req.Ref1, cli.Params.AuthKey)

	//返回值会放到这里
	var result One2PayGenQRCodeResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(cli.Params.PartnerCode, cli.Params.AuthKey, cli.Params.Channel, cli.Params.Device)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//print log
	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#12pay#deposit->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	//-----------错误处理------------------------
	if resp.StatusCode() != 201 {
		if result.Error != "" {
			return nil, errors.New(result.Error)
		}
		return nil, errors.New(result.RespMsg)
	}

	return &result, err
}

/*
curl --location --request POST 'https://api.1-2-pay.com/v1/create-qr-code' \
--header 'channel: WEB' \
--header 'location;' \
--header 'device: WEB' \
--header 'Content-Type: application/json' \
--header 'partner_code: TCU' \
--header 'authorization:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ...' \
--data-raw '{
"amount": 1.012,
"ref1": "20231004142403",
"ref2": "",
"ref3": "",
"ref4": ""
}'
*/
