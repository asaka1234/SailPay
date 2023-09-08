package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 计算签名
func GenSign(params map[string]string, privateSecret string) string {
	// 对请求参数按照字母顺序进行排序并组合
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var signature string = ""
	i := 0
	for _, k := range keys {
		v := params[k]
		if v == "" {
			//空的不参与
			continue
		}
		if i != 0 {
			signature += "&"
		}
		signature += k
		signature += "="
		signature += v
		i++
	}

	signature = fmt.Sprintf("%s&key=%s", signature, privateSecret)

	fmt.Printf("%s\n", signature)

	//md5运算.结果大写
	h := md5.New()
	h.Write([]byte(signature))
	result := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	//amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.google.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.google.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH
	//amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.google.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.google.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH
	//res := md5.Sum([]byte(signature))
	//result := strings.ToUpper(fmt.Sprintf("%x", res))
	return result
}

/*
func GenSignURL(host string) string {

	sign.GenSign()
	query := fmt.Sprintf("time=%v&nonce_str=%v&sign=%v", sign.Time, sign.NonceStr, sign.Sign)
	return host + "?" + query
}

*/
