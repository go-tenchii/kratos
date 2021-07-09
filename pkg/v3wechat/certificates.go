package v3wechat

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
)

func (c *Client) UpdateCertificates() (err error) {
	var b []byte
	var res string
	b, err = c.updateCertificateRequest("GET", certificatesUrl, nil)
	if err != nil {
		return
	}
	rsp := &CertificateRep{}
	if err = json.Unmarshal(b, rsp); err != nil {
		return
	}

	mi := c.GetMchInfo()
	var newestData *CertificateRspData
	for _, v := range rsp.Data {
		if newestData == nil {
			newestData = &v
		} else if newestData.EffectiveTime < v.EffectiveTime {
			newestData = &v
		}
	}

	if newestData == nil {
		return errors.New("证书下载失败")
	}

	res, err = DecryptAES256GCM(mi.ApiKey, newestData.EncryptCertificate.AssociatedData, newestData.EncryptCertificate.Nonce, newestData.EncryptCertificate.Ciphertext)
	mi.WeChatSerialNo = newestData.SerialNo
	mi.Certificate = res
	c.UpdateMchInfo(mi)
	c.updateMchInfoFunc(c)
	return
}

func LoadCertificate(certificateStr string) (certificate *x509.Certificate, err error) {
	block, _ := pem.Decode([]byte(certificateStr))
	if block == nil {
		return nil, fmt.Errorf("decode certificate err")
	}
	certificate, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse certificate err:%s", err.Error())
	}
	return certificate, nil
}
