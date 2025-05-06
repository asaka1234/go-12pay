package go_12pay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 充值的回调通知处理, 将之集成到http server里
func DepositNotify(r *http.Request, processor func(CallbackRequest) error) error {

	//1. 获取返回
	var req CallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req) // 解析响应体为JSON格式
	if err != nil {
		fmt.Println("解析JSON失败：", err)
		return err
	}

	//实际处理
	return processor(req)
}
