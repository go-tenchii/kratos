package v3wechat

import (
	"encoding/json"
	"strings"
)

func (c *Client)PayJsapi(req *PayJsapiReq) (rsp *PayJsapiRsp, err error) {
	var b []byte
	if b, err = json.Marshal(req); err != nil {
		return
	}
	if b, err = c.request("POST", payTransactionsJsapiUrl, b) ; err != nil {
		return
	}
	rsp = &PayJsapiRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

func (c *Client)PayJsapiOutTradeNoQuery(outTradeNo string) (rsp *PayJsapiOutTradeNoRsp, err error) {
	var b []byte
	url := strings.Replace(payOutTradeNoQueryUrl,"{out_trade_no}",outTradeNo,-1)
	url += "?mchid=" + c.mchId
	if b, err = c.request("GET", url, nil) ; err != nil {
		return
	}
	rsp = &PayJsapiOutTradeNoRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}