package sign

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	//代签名query参数
	signMap := map[string]string{
		"platId":     "1000",
		"mchOrderNo": "P0123456789101",
		"amount":     "10000",
		"clientIp":   "192.168.0.111",
		"returnUrl":  "https://www.google.com",
		"notifyUrl":  "https://www.google.com",
		"reqTime":    "20190723141000",
		"version":    "1.0",
	}
	result := GenSign2(signMap, "EWEFD123RGSRETYDFNGFGFGSHDFGH")
	fmt.Printf("%s\n", result)

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
