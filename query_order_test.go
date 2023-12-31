package sailpay_client

import (
	"fmt"
	"testing"
)

// 下单
func TestQueryOrder(t *testing.T) {
	orderId := "exl2ink1a07281"

	//-----------------------------------

	client := NewSailPayClient(MchNo, AppId, PrivateSecret)
	isSucceed, response := client.GetOrderInfo(QueryOrderRequest{
		MchOrderNo: orderId,
	})
	fmt.Printf("result=%v\nresp=%v+\n", isSucceed, response)
}
