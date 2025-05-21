package utils

import (
	"fmt"
	"github.com/spf13/cast"
)

// ------计算一个签名---
func GenSign(amount float64, ref1 string, authKey string) string {
	//把必填的amount和ref1和authKey做一个MD5值,截取位18位后放到ref4里. 这样联合校验
	signRawStr := fmt.Sprintf("amount=%s&ref1=%s&auth=%s", cast.ToString(amount), ref1, authKey)
	signMd5 := GetMD5([]byte(signRawStr))
	sign := signMd5[0:18] //截取18个

	return sign
}

func VerifySign(amount float64, ref1 string, authKey string, sign string) bool {
	signSelf := GenSign(amount, ref1, authKey)
	return signSelf == signSelf
}
