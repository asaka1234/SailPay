package globepay_client

import (
	"fmt"
	"testing"
)

func TestGenMinipOrder(t *testing.T) {
	orderID := "12223344345"

	client := NewGlobePayClient(PartnerCode, CredentialCode)
	isSucceed, response := client.GenMinipOrder(orderID, MinipOrderRequest{
		Description: "desc",                 //订单标题（最大长度128字符，超出自动截取）
		Price:       1,                      //金额，单位为货币最小单位，例如使用100表示GBP1.00
		Currency:    "CNY",                  //币种代码 默认值: GBP,允许值: GBP, CNY
		NotifyURL:   "http://www.baidu.com", //支付通知url，详见支付通知api，不填则不会推送支付通知
		Operator:    "操作员",                  //操作人员标识
		Appid:       AppID,                  //小程序appid
		CustomerID:  CustomerID,             //小程序openid
	})
	fmt.Printf("%v %v", isSucceed, response)
}

func TestGenNewMinipOrder(t *testing.T) {
	orderID := "12223344345"

	client := NewGlobePayClient(PartnerCode, CredentialCode)
	isSucceed, response, err := client.GenNewMinipOrder(orderID, MinipOrderRequest{
		Description: "desc",                 //订单标题（最大长度128字符，超出自动截取）
		Price:       1,                      //金额，单位为货币最小单位，例如使用100表示GBP1.00
		Currency:    "CNY",                  //币种代码 默认值: GBP,允许值: GBP, CNY
		NotifyURL:   "http://www.baidu.com", //支付通知url，详见支付通知api，不填则不会推送支付通知
		Operator:    "操作员",                  //操作人员标识
		Appid:       AppID,                  //小程序appid
		CustomerID:  CustomerID,             //小程序openid
	})
	fmt.Printf("%v %v %v", isSucceed, response, err)
}
