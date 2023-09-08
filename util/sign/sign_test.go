package sign

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	//代签名query参数
	signMap2 := map[string]interface{}{
		"platId":     "1000",
		"mchOrderNo": "P0123456789101",
		"amount":     "10000",
		"clientIp":   "192.168.0.111",
		"returnUrl":  "https://www.google.com",
		"notifyUrl":  "https://www.google.com",
		"reqTime":    "20190723141000",
		"version":    "1.0",
	}
	result2 := GenSign(signMap2, "EWEFD123RGSRETYDFNGFGFGSHDFGH")
	fmt.Printf("%s\n", result2)
}

func TestSign2(t *testing.T) {

	//代签名query参数
	signMap2 := map[string]interface{}{
		"amount":     5,
		"body":       "GoodsDesc",
		"clientIp":   "192.166.1.132",
		"createdAt":  1622016572190,
		"currency":   "inr",
		"extParam":   "{}",
		"ifCode":     "sailpay",
		"mchNo":      "M1621873433953",
		"appId":      "60cc09bce4b0f1c0b83761c9",
		"mchOrderNo": "mho1621934803068",
		"payOrderId": "20210525172643357010",
		"state":      3,
		"subject":    "GOODS",
		"wayCode":    "SAIL_CASHIER",
		"sign":       "C380BEC2BFD727A4B6845133519F3AD6",
	}
	result2 := GenSign(signMap2, "EWEFD123RGSRETYDFNGFGFGSHDFGH")
	fmt.Printf("%s\n", result2)
}
