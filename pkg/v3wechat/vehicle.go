package v3wechat

import (
	"encoding/json"
	"strings"
)

//查询车牌服务开通信息
func (c *Client) VehicleParkingServicesFind(req *VehicleParkingServiceFindReq) (rsp *VehicleParkingServiceFindRsp, err error) {
	var b []byte
	url := vehicleParkingServiceFindUrl
	ret := Struct2mapJson(*req)
	url += "?"
	for k, v := range ret {
		url += k + "=" + v + "&"
	}
	url = url[:len(url)-1]
	b, err = c.request("GET", url, b)
	if err != nil {
		return
	}
	rsp = &VehicleParkingServiceFindRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

//创建停车入场
func (c *Client) VehicleParkingParkings(req *VehicleParkingParkingsReq) (rsp *VehicleParkingParkingsRsp, err error) {
	var b []byte
	url := vehicleParkingParkingsUrl
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	b, err = c.request("POST", url, b)
	if err != nil {
		return
	}
	rsp = &VehicleParkingParkingsRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

//扣费受理
func (c *Client) VehicleTransactionParking(req *VehicleTransactionsParkingReq) (rsp *VehicleTransactionsParkingRsp, err error) {
	var b []byte
	url := vehicleTransactionsParkingUrl
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	b, err = c.request("POST", url, b)
	if err != nil {
		return
	}
	rsp = &VehicleTransactionsParkingRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}

//查询订单
func (c *Client) VehicleTransactionsOutTradeNo(outTradeNo string, subMchId string) (rsp *VehicleTransactionsParkingRsp, err error) {
	var b []byte
	url := strings.Replace(vehicleTransactionsOutTradeNoUrl, "{out_trade_no}", outTradeNo, -1)
	if subMchId != "" {
		url += "?sub_mchid=" + subMchId
	}
	b, err = c.request("GET", url, nil)
	if err != nil {
		return
	}
	rsp = &VehicleTransactionsParkingRsp{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}
	return
}
