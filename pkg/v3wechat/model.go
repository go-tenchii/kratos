package v3wechat

import "fmt"

type MchInfo struct {
	MchId          string
	PublicKey      string
	PrivateKey     string
	SerialNo       string
	WeChatSerialNo string
	ApiKey         string
	Certificate    string
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

type VehicleParkingServiceFindReq struct {
	Appid       string `json:"appid"`
	SubMchId    string `json:"sub_mchid"`
	PlateNumber string `json:"plate_number"`
	PlateColor  string `json:"plate_color"`
	Openid      string `json:"openid"`
}

type VehicleParkingServiceFindRsp struct {
	PlatformNumber  string `json:"platform_number"`
	PlatformColor   string `json:"platform_color"`
	ServiceOpenTime string `json:"service_open_time"`
	Openid          string `json:"openid"`
	ServiceState    string `json:"service_state"`
}

//plate color车牌颜色，枚举值：
//BLUE：蓝色
//GREEN：绿色
//YELLOW：黄色
//BLACK：黑色
//WHITE：白色
//LIMEGREEN：黄绿色
const PC_BLUE = "BLUE"
const PC_GREEN = "GREEN"
const PC_YELLOW = "YELLOW"
const PC_BLACK = "BLACK"
const PC_WHITE = "WHITE"
const PC_LIMEGREEN = "LIMEGREEN"

type VehicleParkingParkingsReq struct {
	SubMchId     string `json:"sub_mchid"`
	OutParkingNo string `json:"out_parking_no"`
	PlateNumber  string `json:"plate_number"`
	PlateColor   string `json:"plate_color"`
	NotifyUrl    string `json:"notify_url"`
	StartTime    string `json:"start_time"`
	ParkingName  string `json:"parking_name"`
	FreeDuration int    `json:"free_duration"`
}

// state enum
const VS_NORMAL = "NORMAL"
const VS_BLOCKED = "BLOCKED"

type VehicleParkingParkingsRsp struct {
	Id           string `json:"id"`
	OutParkingNo string `json:"out_parking_no"`
	PlateNumber  string `json:"plate_number"`
	PlateColor   string `json:"plate_color"`
	StartTime    string `json:"start_time"`
	ParkingName  string `json:"parking_name"`
	FreeDuration int    `json:"free_duration"`
	State        string `json:"state"`
	BlockReason  string `json:"block_reason"`
}

type VehicleTransactionsParkingReq struct {
	Appid         string                           `json:"appid"`
	SubAppid      string                           `json:"sub_appid,omitempty"`
	SubMchId      string                           `json:"sub_mchid,omitempty"`
	Description   string                           `json:"description"`
	Attach        string                           `json:"attach,omitempty"`
	OutTradeNo    string                           `json:"out_trade_no"`
	TradeScene    string                           `json:"trade_scene"`
	GoodsTag      string                           `json:"goods_tag,omitempty"`
	NotifyUrl     string                           `json:"notify_url"`
	ProfitSharing string                           `json:"profit_sharing,omitempty"`
	Amount        VehicleTransactionsParkingAmount `json:"amount"`
	ParkingInfo   VehicleTransactionsParkingInfo   `json:"parking_info"`
}

type VehicleTransactionsParkingAmount struct {
	Total         int    `json:"total"`
	Currency      string `json:"currency,omitempty"`
	PayerTotal    int    `json:"payer_total,omitempty"`
	DiscountTotal int    `json:"discount_total,omitempty"`
}

type VehicleTransactionsParkingInfo struct {
	ParkingId        string `json:"parking_id"`
	PlateNumber      string `json:"plate_number"`
	PlateColor       string `json:"plate_color"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	ParkingName      string `json:"parking_name"`
	ChargingDuration int    `json:"charging_duration"`
	DeviceId         string `json:"device_id"`
}

type VehicleTransactionParkingPayer struct {
	Openid    string `json:"openid"`
	SubOpenid string `json:"sub_openid"`
}

type VehicleTransactionsParkingPromotionDetail struct {
	CouponId            string `json:"coupon_id"`
	Name                string `json:"name"`
	Scope               string `json:"scope"`
	Type                string `json:"type"`
	StockId             string `json:"stock_id"`
	Amount              int    `json:"amount"`
	WeChatPayContribute int    `json:"wechatpay_contribute"`
	MerchantContribute  int    `json:"merchant_contribute"`
	OtherContribute     int    `json:"other_contribute"`
	Currency            string `json:"currency"`
}

//user_repaid 枚举值：
//Y：用户已还款
//N：用户未还款
const UR_Y = "Y"
const UR_N = "N"

//trade state 枚举值：
//SUCCESS：支付成功
//ACCEPTED：已接收，等待扣款
//PAY_FAIL：支付失败(其他原因，如银行返回失败)
//REFUND：转入退款
const TS_SUCCESS = "SUCCESS"
const TS_ACCEPTED = "ACCEPTED"
const TS_PAY_FAIL = "PAY_FAIL"
const TS_REFUND = "REFUND"

type VehicleTransactionsParkingRsp struct {
	Appid                 string                                    `json:"appid"`
	SubAppid              string                                    `json:"sub_appid"`
	SubMchId              string                                    `json:"sub_mchid"`
	SpMchId               string                                    `json:"sp_mchid"`
	Description           string                                    `json:"description"`
	CreateTime            string                                    `json:"create_time"`
	OutTradeNo            string                                    `json:"out_trade_no"`
	TradeState            string                                    `json:"trade_state"`
	TradeStateDescription string                                    `json:"trade_state_description"`
	SuccessTime           string                                    `json:"success_time"`
	BankType              string                                    `json:"bank_type"`
	UserRepaid            string                                    `json:"user_repaid"`
	Attach                string                                    `json:"attach"`
	TradeScene            string                                    `json:"trade_scene"`
	ParkingInfo           VehicleTransactionsParkingInfo            `json:"parking_info"`
	Payer                 VehicleTransactionParkingPayer            `json:"payer"`
	Amount                VehicleTransactionsParkingAmount          `json:"amount"`
	PromotionDetail       VehicleTransactionsParkingPromotionDetail `json:"promotion_detail"`
}

type CertificateRep struct {
	Data []CertificateRspData `json:"data"`
}

type CertificateRspData struct {
	SerialNo           string                           `json:"serial_no"`
	EffectiveTime      string                           `json:"effective_time"`
	ExpireTime         string                           `json:"expire_time"`
	EncryptCertificate CertificateRepEncryptCertificate `json:"encrypt_certificate"`
}

type CertificateRepEncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}
