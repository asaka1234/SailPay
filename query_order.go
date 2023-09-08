package sailpay_client

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"logtech.com/pay/SailPay/util/sign"
)

/*
查询订单详情
*/
func (client *SailPayClient) GetOrderInfo(request QueryOrderRequest) (bool, QueryOrderResponse) {

	var urlResp QueryOrderResponse

	url := QUERY_ORDER_URL

	//计算签名
	rawParams := structs.Map(request)
	signVal := sign.GenSign(rawParams, PrivateSecret)
	request.Sign = signVal

	//构造请求body
	paramJSON, _ := json.Marshal(request)
	paramStr := string(paramJSON)

	//发送请求
	_, _, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		return false, QueryOrderResponse{}
	} else {
		return true, urlResp
	}
}
