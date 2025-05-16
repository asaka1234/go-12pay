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
	Error    string `json:"error"`     //如果返回错误，则有该字段
	RespCode int    `json:"resp_code"` //201是正确
	RespMsg  string `json:"resp_msg"`
	Data     struct {
		Method            string `json:"method"`
		Channel           string `json:"channel"`
		Ref1              string `json:"ref1"`
		Ref2              string `json:"ref2"`
		Ref3              string `json:"ref3"`
		Ref4              string `json:"ref4"`
		Amount            string `json:"amount"`
		Currency          string `json:"currency"`
		Location          string `json:"location"`
		Device            string `json:"device"`
		PartnerCode       string `json:"partner_code"`
		CodeType          string `json:"code_type"`
		CodeImage         string `json:"code_image"`
		CodeURL           string `json:"code_url"`
		TransID           string `json:"trans_id"`
		WalletCode        string `json:"wallet_code"`
		One2payRef        string `json:"one2pay_ref"`
		StoreID           string `json:"store_id"`
		TerminalID        string `json:"terminal_id"`
		MobileNo          string `json:"mobile_no"`
		ExpiredDate       string `json:"expired_date"`
		CreatedDate       string `json:"created_date"`
		CodeImageTemplate string `json:"code_image_template"`
	} `json:"data"`
}

// ----------deposit callback-------------------------
// Notice: 12pay是没有任何callback的验签逻辑的, 所以需要自己搞. 一般都是借助ref字段实现

type One2PayDepositBackReq struct {
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

type One2PayDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
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
