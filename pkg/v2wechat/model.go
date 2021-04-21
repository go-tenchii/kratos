package v2wechat

import "encoding/xml"

type MchInfo struct {
	MchId string
	PublicKey string
	PrivateKey string
	ApiKey string
}

type SceneInfo struct {
	Address string `xml:"address,omitempty"`
	AreaCode string `xml:"area_code,omitempty"`
	Id string `xml:"id"`
	Name string `xml:"name,omitempty"`
}

type UnifiedOrderReq struct {
	Appid string `xml:"appid"`
	Attach string `xml:"attach,omitempty"`
	Body string `xml:"body"`
	Detail string `xml:"detail,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty"`
	GoodsTag string `xml:"goods_tag,omitempty"`
	FeeType string `xml:"fee_type,omitempty"`
	LimitPay string `xml:"limit_pay,omitempty"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	NotifyUrl string `xml:"notify_url"`
	Openid string `xml:"openid,omitempty"`
	OutTradeNo string `xml:"out_trade_no"`
	ProductId string `xml:"product_id,omitempty"`
	ProfitSharing string `xml:"profit_sharing,omitempty"`
	Receipt string `xml:"receipt,omitempty"`
	SceneInfo SceneInfo `xml:"scene_info,omitempty"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type,omitempty"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	TimeExpire string `xml:"time_expire,omitempty"`
	TimeStart string `xml:"time_start,omitempty"`
	TotalFee int `xml:"total_fee"`
	TradeType string `xml:"trade_type"`
}

type OrderQueryReq struct {
	Appid         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type"`
	TransactionId string `xml:"transaction_id,omitempty"`
}

type PayNotifyInfo struct {
	Appid string `xml:"appid"`
	Attach string `xml:"attach,omitempty"`
	BankType string `xml:"bank_type"`
	CashFee int `xml:"cash_fee"`
	CashFeeType string `xml:"cash_fee_type,omitempty"`
	CouponCount int `xml:"coupon_count,omitempty"`
	CouponFee  int `xml:"coupon_fee,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty"`
	ErrCode string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
	FeeType string `xml:"fee_type,omitempty"`
	IsSubscribe string `xml:"is_subscribe"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	OutTradeNo string `xml:"out_trade_no"`
	Openid string `xml:"openid"`
	ResultCode string `xml:"result_code,omitempty"`
	SettlementTotalFee int `xml:"settlement_total_fee,omitempty"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type"`
	TotalFee int `xml:"total_fee"`
	TimeEnd string `xml:"time_end"`
	TradeType string `xml:"trade_type"`
	TransactionId string `xml:"transaction_id"`
}

type PayNotifyRsp struct {
	xml.Name
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg,omitempty"`
}