package v3wechat

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	rand2 "math/rand"
	"time"
)

func Sign(method string, url string, timestamp string, nonce string, body []byte, key string) (sign string, err error)  {
	sign = ""
	sign += method + "\n"
	sign += url + "\n"
	sign += timestamp + "\n"
	sign += nonce + "\n"
	if len(body) != 0 {
		sign += string(body) + "\n"
	} else {
		sign += "\n"
	}
	return PrivateKeySignSHA256(sign, key)
}

func VerifySign(body, timestamp, nonceStr, sign, publicKey string) (err error)  {
	bf := bytes.Buffer{}
	bf.WriteString(timestamp + "\n")
	bf.WriteString(nonceStr + "\n")
	bf.WriteString(body)
	return RsaVerySignWithSHA256Base64(bf.String(), sign, publicKey)
}

func PrivateKeySignSHA256(signData string, key string) (sign string, err error) {
	var priv interface{}
	sign = ""
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		err = errors.New("pem.Decode err")
		return
	}
	priv, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	sign, err = SHA256BAS64(signData, priv.(*rsa.PrivateKey))
	if err != nil {
		return
	}
	return
}

func SHA256BAS64(signstr string, key *rsa.PrivateKey) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(signstr))
	hased := hasher.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hased)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := []byte(str)
	result := []byte{}
	r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func RsaVerySignWithSHA256Base64(originalData, signData string, key string) (err error) {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, h.Sum(nil), sign)
}
