package v3wechat

import (
	"encoding/json"
	"strings"
)

func (c *Client)SendCouponStock(req *SendCouponStockReq) (rsp *SendCouponStockRsp, err error) {
	var b []byte
	sendCouponStockUrl := strings.Replace(sendCouponStockUrl, "{openid}", req.Openid, -1)
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

func (c *Client)QueryUserCoupon(req *QueryUserCouponReq) (rsp *QueryUserCouponRsp, err error) {
	var b []byte
	queryUserCouponUrl := strings.Replace(userCouponQueryUrl, "{openid}", req.Openid, -1)
	queryUserCouponUrl = strings.Replace(userCouponQueryUrl, "{coupon_id}", req.CouponId, -1)
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	b, err = c.request("POST", queryUserCouponUrl, b)
	if err != nil {
		return
	}
	rsp = &QueryUserCouponRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}