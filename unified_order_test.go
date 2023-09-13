package sailpay_client

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 下单
func TestPlaceOrder(t *testing.T) {
	orderId := "exop20230908002"

	params := map[string]string{
		"firstname": "cy",
		"lastname":  "harper",
		"city":      "guangzhou",
		"phone":     "4401000001",
		"email":     "ck789@gmail.com",
		"country":   "IN",
		"address":   "baiyun district",
		"state":     "mh",
		"postcode":  "232001",
	}
	res, _ := json.Marshal(params)

	//-----------------------------------

	client := NewSailPayClient(MchNo, AppId, PrivateSecret)
	isSucceed, response := client.PlaceUnifiedOrder(UnifiedOrderRequest{
		MchOrderNo: orderId,
		WayCode:    "SAIL_CASHIER",
		Amount:     1000,
		Currency:   "inr",
		Subject:    "toys",
		Body:       "toysDesc",
		NotifyUrl:  "https://www.jpdb001.com/notifyUrl",
		ReturnUrl:  "https://www.jpdb001.com/returnUrl",
		ExtParam:   string(res),
		//可选
		ExpiredTime: 3600,
		//ClientIp:    "192.166.1.132",
		//ChannelExtra: "{\"authCode\":\"280812820366966512\"}",
		DivisionMode: 0,
	})
	fmt.Printf("result=%v\nresp=%v+\n", isSucceed, response)
}
