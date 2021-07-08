package v3wechat

const (
	baseUrl = "https://api.mch.weixin.qq.com"

	sendCouponStockUrl  = "/v3/marketing/favor/users/{openid}/coupons"
	userCouponQueryUrl  = "/v3/marketing/favor/users/{openid}/coupons/{coupon_id}"
	userCouponsQueryUrl = "/v3/marketing/favor/users/{openid}/coupons"

	payTransactionsJsapiUrl = "v3/pay/transactions/jsapi"
	payOutTradeNoQueryUrl   = "v3/pay/transactions/out-trade-no/{out_trade_no}"
)
