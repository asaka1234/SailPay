package sailpay_client

import (
	"fmt"
	"testing"
)

// 查询settle
func TestSettleListOrder(t *testing.T) {

	client := NewSailPayClient(MchNo, AppId, PrivateSecret)
	isSucceed, response := client.QuerySettleList(QuerySettleListRequest{
		Currency:     "BRL",
		CreatedStart: "2023-10-07 12:00:00",
		CreatedEnd:   "2023-10-13 12:00:00",
		SettleState:  "1",
	})
	fmt.Printf("result=%v\nresp=%v+\n", isSucceed, response)
}
