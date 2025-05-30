package go_12pay

import (
	"errors"
	"github.com/asaka1234/go-12pay/utils"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req One2PayDepositBackReq, processor func(One2PayDepositBackReq) error) error {
	//验证签名
	verifySucceed := utils.VerifySign(req.Amount, req.Ref1, cli.Params.AuthKey, req.Ref4)
	if !verifySucceed {
		return errors.New("verify failed!")
	}
	//看看是否成功
	if req.RespCode != int(One2payDepositStatusSuccess) {
		//说明失败
		return errors.New("resp_msg is " + req.RespMsg)
	}

	//开始处理
	return processor(req)
}
