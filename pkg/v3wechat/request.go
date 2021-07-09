package v3wechat

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-tenchii/kratos/pkg/log"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) request(method string, route string, body []byte) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	dc := http.DefaultClient

	req, err = http.NewRequest(method, baseUrl+route, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	mi := c.GetMchInfo()
	timestamp := fmt.Sprint(time.Now().Unix())
	nonceStr := GetRandomString(32)

	sign, err := Sign(method, route, timestamp, nonceStr, body, mi.PrivateKey)
	if err != nil {
		return nil, err
	}

	auth := "WECHATPAY2-SHA256-RSA2048 " +
		"mchid=\"" + mi.MchId +
		"\",nonce_str=\"" + nonceStr +
		"\",signature=\"" + sign +
		"\",timestamp=\"" + timestamp +
		"\",serial_no=\"" + mi.SerialNo + "\""
	req.Header.Set("User-Agent", "v3wechat-go-tenchii")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", auth)

	log.Info("nonceStr: %s, route: %s, req: %s", nonceStr, route, string(body))
	resp, err = dc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Info("nonceStr: %s, request_id: %s, route: %s, rsp: %s", nonceStr, resp.Header.Get("Request-Id"), route, string(rspBody))
	if resp.StatusCode != http.StatusOK {
		be := &RequestErr{}
		if err = json.Unmarshal(rspBody, be); err != nil {
			return nil, err
		}
		return nil, be
	}

	serialNo := resp.Header.Get("Wechatpay-Serial")
	if serialNo != mi.WeChatSerialNo { //不一致
		if err = c.UpdateCertificates(); err != nil {
			log.Error("UpdateCertificates err: %s", err)
			return rspBody, nil
		} else {
			mi = c.GetMchInfo()
		}
	}

	wpTimestamp := resp.Header.Get("Wechatpay-Timestamp")
	wpNonceStr := resp.Header.Get("Wechatpay-Nonce")
	wpSign := resp.Header.Get("Wechatpay-Signature")
	if wpSign == "" || wpTimestamp == "" || wpNonceStr == "" {
		return nil, errors.New("伪造或被篡改的应答")
	}

	cert, _ := LoadCertificate(mi.Certificate)
	if err = VerifySign(string(rspBody), wpTimestamp, wpNonceStr, wpSign, cert.PublicKey.(*rsa.PublicKey)); err != nil {
		return nil, errors.New("签名验证错误")
	}

	return rspBody, nil
}

func (c *Client) updateCertificateRequest(method string, route string, body []byte) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	dc := http.DefaultClient

	req, err = http.NewRequest(method, baseUrl+route, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	mi := c.GetMchInfo()
	timestamp := fmt.Sprint(time.Now().Unix())
	nonceStr := GetRandomString(32)

	sign, err := Sign(method, route, timestamp, nonceStr, body, mi.PrivateKey)
	if err != nil {
		return nil, err
	}

	auth := "WECHATPAY2-SHA256-RSA2048 " +
		"mchid=\"" + mi.MchId +
		"\",nonce_str=\"" + nonceStr +
		"\",signature=\"" + sign +
		"\",timestamp=\"" + timestamp +
		"\",serial_no=\"" + mi.SerialNo + "\""
	req.Header.Set("User-Agent", "v3wechat-go-tenchii")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", auth)

	log.Info("nonceStr: %s, route: %s, req: %s", nonceStr, route, string(body))
	resp, err = dc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		be := &RequestErr{}
		if err = json.Unmarshal(rspBody, be); err != nil {
			return nil, err
		}
		return nil, be
	}
	log.Info("nonceStr: %s, request_id: %s, route: %s, rsp: %s", nonceStr, resp.Header.Get("Request-Id"), route, string(rspBody))
	return rspBody, nil
}
