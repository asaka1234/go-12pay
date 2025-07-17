package go_12pay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//cc := One2payDepositStatusSuccess

	//构造client
	cli := NewClient(vLog, &One2PayInitParams{PARTNER_CODE, AUTH_KEY, DEVICE, CHANNEL, DEPOSIT_URL, PAYOUT_URL})

	//发请求
	cli.SetDebugModel(true)
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() One2PayWithdrawRequest {
	return One2PayWithdrawRequest{
		BankAcc:       "123456789012",
		BankCode:      "004",
		BankName:      "KASIKORN BANK",
		AccountName:   "Manop Tangngam",
		Amount:        1000.50,
		MobileNo:      "0805933181",
		TransactionBy: "Jack Developer",
		Ref1:          "12345",
	}
}
