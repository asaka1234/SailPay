package sailpay_client

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"logtech.com/pay/SailPay/util/sign"
)

// 生成账单
/*
	orderID 商户内部订单id，要求同一商户唯一
*/
func (client *SailPayClient) PlaceUnifiedOrder(request UnifiedOrderRequest) (bool, UnifiedOrderResponse) {

	var urlResp UnifiedOrderResponse

	url := UNIFIEDORDER_URL

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
		return false, UnifiedOrderResponse{}
	} else {
		return true, urlResp
	}
}
