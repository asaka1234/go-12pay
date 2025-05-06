package go_12pay

import "time"

// ----------generate qrcode-------------------------

// 生成二维码
type GenQRCodeRequest struct {
	Amount float64 `json:"amount"` //金额
	Ref1   string  `json:"ref1"`   //可以放业务自己的orderNo  同时可以加上一些验签
	Ref2   string  `json:"ref2"`
	Ref3   string  `json:"ref3"`
	Ref4   string  `json:"ref4"`
}

type GenQRCodeResponse struct {
	RespCode int    `json:"resp_code"` //201是正确
	RespMsg  string `json:"resp_msg"`
	Data     struct {
		PartnerCode string    `json:"partner_code"`
		Method      string    `json:"method"`
		CodeType    string    `json:"code_type"`
		CodeImage   string    `json:"code_image"`
		CodeUrl     string    `json:"code_url"` //是付款二维码的url
		TransId     string    `json:"trans_id"`
		Channel     string    `json:"channel"`
		WalletCode  string    `json:"wallet_code"`
		Ref1        string    `json:"ref1"`
		Ref2        string    `json:"ref2"`
		Ref3        string    `json:"ref3"`
		Ref4        string    `json:"ref4"`
		One2PayRef  string    `json:"one2pay_ref"`
		Amount      int       `json:"amount"`
		Currency    string    `json:"currency"`
		StoreId     string    `json:"store_id"`
		TerminalId  string    `json:"terminal_id"`
		MobileNo    string    `json:"mobile_no"`
		Location    string    `json:"location"`
		Device      string    `json:"device"`
		ExpiredDate string    `json:"expired_date"`
		CreatedDate time.Time `json:"created_date"`
	} `json:"data"`
}

// ----------deposit callback-------------------------
// Notice: 12pay是没有任何callback的验签逻辑的, 所以需要自己搞. 一般都是借助ref字段实现

type CallbackRequest struct {
	RespCode   int     `json:"resp_code"` //200是成功
	RespMsg    string  `json:"resp_msg"`
	Command    string  `json:"command"`
	BankRef    string  `json:"bank_ref"`
	TranxId    string  `json:"tranx_id"`
	One2PayRef string  `json:"one2pay_ref"`
	Datetime   string  `json:"datetime"`
	EffDate    string  `json:"effdate"`
	Amount     float64 `json:"amount"`
	CusName    string  `json:"cusname"`
	Ref1       string  `json:"ref1"`
	Ref2       string  `json:"ref2"`
	Ref3       string  `json:"ref3"`
	Ref4       string  `json:"ref4"`
	TransId    string  `json:"trans_id"`
}

// ----------withdraw-------------------------

type WithdrawRequest struct {
	BankAcc       string  `json:"bankacc"`        //required
	BankCode      string  `json:"bankcode"`       //required
	BankName      string  `json:"bankname"`       //required
	AccName       string  `json:"accname"`        //required
	Amount        float64 `json:"amount"`         //required
	MobileNo      string  `json:"mobileno"`       //required
	TransactionBy string  `json:"transaction_by"` //required
	Ref1          string  `json:"ref1"`           //required
}

type WithdrawResponse struct {
	Status              int       `json:"status"` //Success Case Status 1000 only
	Message             string    `json:"message"`
	Ref1                string    `json:"ref1"`
	TransactionId       string    `json:"transaction_id"`
	TransactionDateTime time.Time `json:"transactionDate_time"` //YYY-MM-DD hh:mm:ss
}
