package sailpay_client

import (
	"fmt"
	"testing"
)

// 查询余额
func TestQueryBalance(t *testing.T) {
	//orderId := "exl2ink1a07281"

	//-----------------------------------

	client := NewSailPayClient(MchNo, AppId, PrivateSecret)
	isSucceed, response := client.GetBalanceInfo(QueryBalanceRequest{
		Currency: "BRL",
	})
	fmt.Printf("result=%v\nresp=%v+\n", isSucceed, response)
}
