package go_12pay

import "time"

type One2PayInitParams struct {
	PartnerCode string `json:"partnerCode" mapstructure:"partnerCode" config:"partnerCode"` //商户号
	AuthKey     string `json:"authKey" mapstructure:"authKey" config:"authKey"`             //accessKey
	Device      string `json:"device" mapstructure:"device" config:"device"`
	Channel     string `json:"channel" mapstructure:"channel" config:"channel"`

	DepositUrl  string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"`    //充值url
	WithdrawUrl string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"` //退款url
}

// ----------generate qrcode-------------------------

// 生成二维码
type One2PayGenQRCodeRequest struct {
	Amount float64 `json:"amount"` //must 金额(不需要做单位转换) 只是THB泰铢
	Ref1   string  `json:"ref1"`   //must 放业务自己的orderNo (只能是数字/字母) The ref1 use bank format can’t more than 18 digit
	//option
	Ref2 string `json:"ref2"`
	Ref3 string `json:"ref3"` //放customer name
	Ref4 string `json:"ref4"` // 这个用来做签名, 是amount/ref1/authkey的一个md5签名的截断值(18位)
}

type One2PayGenQRCodeResponse struct {
	Error    string `json:"error"`     //如果返回错误，则有该字段
	RespCode int    `json:"resp_code"` //201是正确
	RespMsg  string `json:"resp_msg"`
	Data     struct {
		Method            string  `json:"method"`
		Channel           string  `json:"channel"`
		Ref1              string  `json:"ref1"`
		Ref2              string  `json:"ref2"`
		Ref3              string  `json:"ref3"`
		Ref4              string  `json:"ref4"` // 这个用来做签名, 是amount/ref1/authkey的一个md5签名的截断值(18位)
		Amount            float64 `json:"amount"`
		Currency          string  `json:"currency"` //THB
		Location          string  `json:"location"`
		Device            string  `json:"device"`
		PartnerCode       string  `json:"partner_code"`
		CodeType          string  `json:"code_type"` //ThaiQRCode
		CodeImage         string  `json:"code_image"`
		CodeURL           string  `json:"code_url"`
		TransID           string  `json:"trans_id"`
		WalletCode        string  `json:"wallet_code"`
		One2payRef        string  `json:"one2pay_ref"`
		StoreID           string  `json:"store_id"`
		TerminalID        string  `json:"terminal_id"`
		MobileNo          string  `json:"mobile_no"`
		ExpiredDate       string  `json:"expired_date"`
		CreatedDate       string  `json:"created_date"`
		CodeImageTemplate string  `json:"code_image_template"`
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
	Ref1       string  `json:"ref1"` //放业务自己的orderNo (只能是数字/字母)
	Ref2       string  `json:"ref2"`
	Ref3       string  `json:"ref3"`
	Ref4       string  `json:"ref4"`     // 这个用来做签名, 是amount/ref1/authkey的一个md5签名的截断值(18位)
	TransId    string  `json:"trans_id"` //是psp三方的订单id
}

/*
{
	"resp_code":200,
	"resp_msg":"Success“,
	"command":"Payment",
	"bank_ref":"ITMX 13078496",
	"tranx_id":"5oaY6TEYxKPBi34yDxM2",
	"one2pay_ref":"ABC220407080726611",
	"datetime":"20231004133121",
	"effdate":"20231004",
	"amount":18658.26,
	"cusname":"นาย จิณณะ แสงฤทธิ์“,
	"ref1":"SPI729131696400892",
	"ref2":“123456",
	"ref3":“abc123-",
	"ref4":"abc123-",
	"trans_id":"2f0cf4e1ceb53c4053837d0860c19620“
}
*/

type One2PayDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// ----------withdraw-------------------------

type One2PayWithdrawRequest struct {
	BankAcc       string  `json:"bankacc"`        //required
	BankCode      string  `json:"bankcode"`       //required
	BankName      string  `json:"bankname"`       //required
	AccountName   string  `json:"accname"`        //required
	Amount        float64 `json:"amount"`         //required
	MobileNo      string  `json:"mobileno"`       //required
	TransactionBy string  `json:"transaction_by"` //required
	Ref1          string  `json:"ref1"`           //required
	Ref2          string  `json:"ref2"`
	Ref3          string  `json:"ref3"`
	Ref4          string  `json:"ref4"`
}

/*
{
	"bankacc":“0652078409",
	"bankcode":"004",
	"bankname":"KASIKORN BANK",
	"accname":"Manop Tangngam",
	"amount":1000.50,
	"mobileno":"0805933181",
	"transaction_by":"Jack Developer",
	"ref1": "123456789012345678“
}
*/

type One2PayWithdrawResponse struct {
	Error               string    `json:"error"`
	Status              int       `json:"status"` //Success Case Status 1000 only
	Message             string    `json:"message"`
	Ref1                string    `json:"ref1"`
	TransactionId       string    `json:"transaction_id"`
	TransactionDateTime time.Time `json:"transactionDate_time"` //YYY-MM-DD hh:mm:ss
}
