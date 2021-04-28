package v2wechat

import (
	"encoding/xml"
	"fmt"
	"github.com/go-tenchii/kratos/pkg/ecode"
	"github.com/go-tenchii/kratos/pkg/log"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client)UnifiedOrder(reqS *UnifiedOrderReq) (rsp StringMap, err error) {
	req := StringMap{}
	req["appid"] = reqS.Appid
	if len(reqS.SubAppid) > 0 {
		req["sub_appid"] = reqS.SubAppid
	}
	if len(reqS.SubMchId) > 0 {
		req["sub_mch_id"] = reqS.SubMchId
	}
	req["mch_id"] = reqS.MchId
	req["nonce_str"] = GetRandomString(32)
	req["body"] = reqS.Body
	req["out_trade_no"] = reqS.OutTradeNo
	req["total_fee"] = fmt.Sprint(reqS.TotalFee)
	req["spbill_create_ip"] = reqS.SpbillCreateIp
	req["notify_url"] = reqS.NotifyUrl
	req["trade_type"] = reqS.TradeType
	if reqS.Openid != "" {
		req["openid"] = reqS.Openid
	}
	if reqS.SubOpenid != "" {
		req["sub_openid"] = reqS.SubOpenid
	}
	if reqS.Attach != "" {
		req["attach"] = reqS.Attach
	}
	req["time_start"] = reqS.TimeStart
	req["time_expire"] = reqS.TimeExpire
	req["sign"] = WeChatPaySignMd5WithStringMap(req, c.getMchInfoFunc(c).ApiKey)
	buf, err := xml.Marshal(req)
	if err != nil {
		return
	}
	log.Info("pay", "Unified Order req : "+string(buf))
	resp, err := http.Post(baseUrl + UnifiedOrderUrl, "text/xml", strings.NewReader(string(buf)))
	if err != nil {
		return
	}
	robots, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = xml.Unmarshal(robots, &rsp)
	if err != nil {
		return
	}
	log.Info("pay", "Unified Order rsp : "+string(robots))
	if rsp["return_code"] != "SUCCESS" {
		log.Error("pay", "unified order fail:"+rsp["return_code"]+" : "+rsp["return_msg"])
		err = ecode.Error(10001,rsp["return_msg"])
		return
	}
	if rsp["result_code"] != "SUCCESS" {
		log.Error("pay", "unified order fail:"+rsp["err_code"]+" : "+rsp["err_code_des"])
		err = ecode.Error(10001,rsp["err_code_des"])
		return
	}
	if WeChatPaySignMd5WithStringMap(rsp, c.getMchInfoFunc(c).ApiKey) != rsp["sign"] {
		err = ecode.Error(10001,"签名错误")
		return
	}
	return
}

func (c *Client)OrderQuery(appid,subAppid,subMchId,outTradeNo string) (rsp StringMap, err error) {
	req := &OrderQueryReq{}
	req.Appid = appid
	req.OutTradeNo = outTradeNo
	req.MchId = c.mchId
	if len(subAppid) > 0 {
		req.SubAppid = subAppid
	}
	if len(subMchId) > 0 {
		req.SubMchId = subMchId
	}
	req.NonceStr = GetRandomString(32)
	req.SignType = "HMAC-SHA256"
	req.Sign = WeChatPaySignHMACSHA256(*req, c.getMchInfoFunc(c).ApiKey)

	v, _ := xml.Marshal(req)
	log.Info("pay", req.OutTradeNo+" send query micro pay req xml: "+string(v))

	var rspWx *http.Response
	var body []byte
	rspWx, err = http.Post(baseUrl + queryOrderUrl, "text/xml", strings.NewReader(string(v)))
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(rspWx.Body)
	defer rspWx.Body.Close()
	if err != nil {
		return
	}

	if err = xml.Unmarshal(body, &rsp); err != nil {
		return
	}
	log.Info("pay", req.OutTradeNo+" send query micro pay rsp xml: "+string(body))
	if rsp["return_code"] != "SUCCESS" {
		log.Error("pay", "unified order fail:"+rsp["return_code"]+" : "+rsp["return_msg"])
		err = ecode.Error(10001,rsp["return_msg"])
		return
	}
	if rsp["result_code"] != "SUCCESS" {
		log.Error("pay", "unified order fail:"+rsp["err_code"]+" : "+rsp["err_code_des"])
		err = ecode.Error(10001,rsp["err_code_des"])
		return
	}
	if WeChatPaySignHMACSHA256ByMap(rsp, c.getMchInfoFunc(c).ApiKey) != rsp["sign"] {
		err = ecode.Error(10001,"签名错误")
		return
	}
	return
}