package sailpay_client

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 下单
func TestPlaceOrder(t *testing.T) {
	orderId := "exl2ink1a07281"

	params := map[string]string{
		"firstname": "Jamer",
		"lastname":  "havi",
		"city":      "Mubai",
		"phone":     "8901000001",
		"email":     "havi@gmail.com",
		"country":   "IN",
		"address":   "test",
		"state":     "mh",
		"postcode":  "23456",
	}
	res, _ := json.Marshal(params)

	//-----------------------------------

	client := NewSailPayClient(MchNo, AppId, PrivateSecret)
	isSucceed, response := client.PlaceUnifiedOrder(UnifiedOrderRequest{
		MchOrderNo: orderId,
		WayCode:    "SAIL_CASHIER",
		Amount:     1000,
		Currency:   "inr",
		Subject:    "GOODS",
		Body:       "GoodsDesc",
		NotifyUrl:  "https://www.yourdomain.com/notifyUrl",
		ReturnUrl:  "https://www.yourdomain.com/returnUrl",
		ExtParam:   string(res),
		//可选
		ExpiredTime: 3600,
		//ClientIp:    "192.166.1.132",
		//ChannelExtra: "{\"authCode\":\"280812820366966512\"}",
		DivisionMode: 1,
	})
	fmt.Printf("result=%v\nresp=%v+\n", isSucceed, response)
}
