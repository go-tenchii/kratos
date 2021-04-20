package v2wechat

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func WeChatPaySignMd5WithStringMap(data map[string]string, apikey string) string {
	context := StringMapJointByDictionaryOrder(data)
	context += "key=" + apikey
	hasher := md5.New()
	hasher.Write([]byte(context))
	sign := hex.EncodeToString(hasher.Sum(nil))
	return strings.ToUpper(sign)
}

func StringMapJointByDictionaryOrder(data map[string]string) (str string) {
	str = ""
	keyArr := []string{}
	for k, _ := range data {
		if k == "sign" {
			continue
		}
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	for _, key := range keyArr {
		str += key + "=" + data[key] + "&"
	}
	return
}

func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func WeChatPaySignHMACSHA256(v interface{}, apikey string) string {
	signstr := ToKeyValueStr(v)
	signstr = signstr + "key=" + apikey
	key := []byte(apikey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signstr))
	sha := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sha)
}

func ToKeyValueStr(v interface{}) string {
	var signstr bytes.Buffer

	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		name := field.Name

		keytemp := field.Tag.Get("xml")
		keymap := strings.Split(keytemp, ",")
		key := keymap[0]
		if reflect.Indirect(vv).FieldByName(name).Type().Kind() == reflect.Int || reflect.Indirect(vv).FieldByName(name).Type().Kind() == reflect.Int64 {
			value := reflect.Indirect(vv).FieldByName(name).Int()
			signstr.WriteString(key + "=" + strconv.Itoa(int(value)) + "&")
		} else if reflect.Indirect(vv).FieldByName(name).Type().Kind() == reflect.Float64 {
			value := reflect.Indirect(vv).FieldByName(name).Float()
			signstr.WriteString(key + "=" + strconv.FormatFloat(value, 'f', -1, 64) + "&")
		} else if reflect.Indirect(vv).FieldByName(name).Type().Kind() == reflect.Uint64 {
			value := reflect.Indirect(vv).FieldByName(name).Uint()
			signstr.WriteString(key + "=" + strconv.FormatUint(value, 10) + "&")
		} else if reflect.Indirect(vv).FieldByName(name).Type().Kind() == reflect.Uint32 {
			value := reflect.Indirect(vv).FieldByName(name).Uint()
			signstr.WriteString(key + "=" + strconv.FormatUint(value, 10) + "&")
		} else {
			value := (reflect.Indirect(vv).FieldByName(name)).String()
			if value != "" && key != "xml" && key != "sign" {
				signstr.WriteString(key + "=" + value + "&")
			}
		}
	}
	//	str := Substr(signstr.String(), 0, len(signstr.String())-1)
	return signstr.String()
}