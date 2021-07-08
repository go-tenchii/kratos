package v3wechat

import "fmt"

type MchInfo struct {
	MchId      string
	PublicKey  string
	PrivateKey string
	SerialNo   string
}

type RequestErr struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Field   string      `json:"field"`
	Value   string      `json:"value"`
	Issue   string      `json:"issue"`
	Detail  interface{} `json:"detail"`
	ReqText string      `json:"req_text"`
}

func (e *RequestErr) Error() string {
	return fmt.Sprintf("code: %s, msg: %s, field: %s, value: %s, issue: %s, detail: %v, reqText: %s", e.Code, e.Message, e.Field, e.Value, e.Issue, e.Detail, e.ReqText)
}

type SendCouponStockReq struct {
	Appid             string `json:"appid"`
	CouponMinimum     uint64 `json:"coupon_minimum,omitempty"`
	CouponValue       uint64 `json:"coupon_value,omitempty"`
	OutRequestNo      string `json:"out_request_no"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockId           string `json:"stock_id"`
}

type SendCouponStockRsp struct {
	CouponId string `json:"coupon_id"`
}

type QueryUserCouponReq struct {
}

type QueryUserCouponRsp struct {
	StockCreateMchId        string                                 `json:"stock_creator_mchid"`
	StockId                 string                                 `json:"stock_id"`
	CouponId                string                                 `json:"coupon_id"`
	CutToMessage            QueryUserCouponCutToMessage            `json:"cut_to_message"`
	CouponName              string                                 `json:"coupon_name"`
	Status                  string                                 `json:"status"`
	Description             string                                 `json:"description"`
	CreateTime              string                                 `json:"create_time"`
	CouponType              string                                 `json:"coupon_type"`
	NoCash                  bool                                   `json:"no_cash"`
	AvailableBeginTime      string                                 `json:"available_begin_time"`
	AvailableEndTime        string                                 `json:"available_end_time"`
	Singleitem              bool                                   `json:"singleitem"`
	NormalCouponInformation QueryUserCouponNormalCouponInformation `json:"normal_coupon_information"`
	ConsumeInformation      QueryUserCouponConsumeInformation      `json:"consume_information"`
	GoodsDetail             QueryUserCouponGoodsDetail             `json:"goods_detail"`
}

type QueryUserCouponCutToMessage struct {
	SinglePriceMax int64 `json:"single_price_max"`
	CutToPrice     int64 `json:"cut_to_price"`
}

type QueryUserCouponNormalCouponInformation struct {
	CouponAmount       uint64 `json:"coupon_amount"`
	TransactionMinimum uint64 `json:"transaction_minimum"`
}

type QueryUserCouponConsumeInformation struct {
	ConsumeTime   string `json:"consume_time"`
	ConsumeMchid  string `json:"consume_mchid"`
	TransactionId string `json:"transaction_id"`
}

type QueryUserCouponGoodsDetail struct {
	GoodsId        string `json:"goods_id"`
	Quantity       uint32 `json:"quantity"`
	Price          uint64 `json:"price"`
	DiscountAmount uint64 `json:"discount_amount"`
}

type PayJsapiReq struct {
	Appid       string             `json:"appid"`
	Mchid       string             `json:"mchid"`
	Description string             `json:"description"`
	OutTradeNo  string             `json:"out_trade_no"`
	TimeExpire  string             `json:"time_expire,omitempty"`
	Attach      string             `json:"attach,omitempty"`
	NotifyUrl   string             `json:"notify_url"`
	GoodsTag    string             `json:"goods_tag,omitempty"`
	Amount      PayJsapiAmountInfo `json:"amount"`
	Payer       PayJsapiPayerInfo  `json:"payer"`
}

type PayJsapiRsp struct {
	PrepayId string `json:"prepay_id"`
}

type PayJsapiAmountInfo struct {
	Total    int    `json:"total"`
	Currency string `json:"currency,omitempty"`
}

type PayJsapiPayerInfo struct {
	Openid string `json:"openid"`
}

type PayJsapiOutTradeNoRsp struct {
	Appid          string            `json:"appid"`
	Mchid          string            `json:"mchid"`
	OutTradeNo     string            `json:"out_trade_no"`
	TransactionId  string            `json:"transaction_id,omitempty"`
	TradeType      string            `json:"trade_type,omitempty"`
	TradeState     string            `json:"trade_state"`
	TradeStateDesc string            `json:"trade_state_desc"`
	BankType       string            `json:"bank_type,omitempty"`
	Attach         string            `json:"attach,omitempty"`
	SuccessTime    string            `json:"success_time,omitempty"`
	Payer          PayJsapiPayerInfo `json:"payer"`
}

type QueryUserCouponsReq struct {
	Appid          string `json:"appid"`
	StockId        string `json:"stock_id,omitempty"`
	Status         string `json:"status,omitempty"`
	CreatorMchId   string `json:"creator_mchid,omitempty"`
	SenderMchId    string `json:"sender_mchid,omitempty"`
	AvailableMchId string `json:"available_mchid,omitempty"`
	Offset         uint32 `json:"offset,omitempty"`
	Limit          uint32 `json:"limit,omitempty"`
}

type QueryUserCouponsRsp struct {
	Data       []QueryUserCouponRsp `json:"data"`
	TotalCount uint32               `json:"total_count"`
	Limit      uint32               `json:"limit"`
	Offset     uint32               `json:"offset"`
}
