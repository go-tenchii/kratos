package v3wechat

import (
	"encoding/json"
	"strings"
)

func (c *Client) SendCouponStock(openid string, req *SendCouponStockReq) (rsp *SendCouponStockRsp, err error) {
	var b []byte
	sendCouponStockUrl := strings.Replace(sendCouponStockUrl, "{openid}", openid, -1)
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	b, err = c.request("POST", sendCouponStockUrl, b)
	if err != nil {
		return
	}
	rsp = &SendCouponStockRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

func (c *Client) QueryUserCoupon(openid string, couponId string, appid string) (rsp *QueryUserCouponRsp, err error) {
	var b []byte
	queryUserCouponUrl := strings.Replace(userCouponQueryUrl, "{openid}", openid, -1)
	queryUserCouponUrl = strings.Replace(queryUserCouponUrl, "{coupon_id}", couponId, -1)
	queryUserCouponUrl += "?appid=" + appid
	b, err = c.request("GET", queryUserCouponUrl, b)
	if err != nil {
		return
	}
	rsp = &QueryUserCouponRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

func (c *Client) QueryUserCoupons(openid string, req *QueryUserCouponsReq) (rsp *QueryUserCouponsRsp, err error) {
	var b []byte
	queryUserCouponsUrl := strings.Replace(userCouponsQueryUrl, "{openid}", openid, -1)
	ret := Struct2mapJson(*req)
	queryUserCouponsUrl += "?"
	for k, v := range ret {
		queryUserCouponsUrl += k + "=" + v + "&"
	}
	queryUserCouponsUrl = queryUserCouponsUrl[:len(queryUserCouponsUrl)-1]
	b, err = c.request("GET", queryUserCouponsUrl, b)
	if err != nil {
		return
	}
	rsp = &QueryUserCouponsRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}
