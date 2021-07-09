package v3wechat

const (
	baseUrl = "https://api.mch.weixin.qq.com"

	certificatesUrl = "/v3/certificates"

	sendCouponStockUrl  = "/v3/marketing/favor/users/{openid}/coupons"
	userCouponQueryUrl  = "/v3/marketing/favor/users/{openid}/coupons/{coupon_id}"
	userCouponsQueryUrl = "/v3/marketing/favor/users/{openid}/coupons"

	vehicleParkingServiceFindUrl     = "/v3/vehicle/parking/services/find"
	vehicleParkingParkingsUrl        = "/v3/vehicle/parking/parkings"
	vehicleTransactionsParkingUrl    = "/v3/vehicle/transactions/parking"
	vehicleTransactionsOutTradeNoUrl = "/v3/vehicle/transactions/out-trade-no/{out_trade_no}"

	payTransactionsJsapiUrl = "v3/pay/transactions/jsapi"
	payOutTradeNoQueryUrl   = "v3/pay/transactions/out-trade-no/{out_trade_no}"
)
