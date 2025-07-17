package go_12pay

import (
	"fmt"
	"testing"
)

func TestGenQRCode(t *testing.T) {
	vlog := VLog{}
	//构造client
	cli := NewClient(vlog, &One2PayInitParams{PARTNER_CODE, AUTH_KEY, DEVICE, CHANNEL, DEPOSIT_URL, PAYOUT_URL})

	//发请求
	resp, err := cli.GenQRCode(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() One2PayGenQRCodeRequest {
	return One2PayGenQRCodeRequest{
		Amount: 700,
		Ref1:   "9438239651111", //Unique customer id in your system. 业务系统里的唯一客户id
		Ref3:   "uname",
	}
}
