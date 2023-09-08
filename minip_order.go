package sailpay_client

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"logtech.com/pay/SailPay/util/sign"
)

// 生成账单
/*
	orderID 商户内部订单id，要求同一商户唯一
*/
func (client *SailPayClient) GenMinipOrder(orderID string, request MinipOrderRequest) (bool, MinipOrderResponse) {

	var urlResp MinipOrderResponse

	url := fmt.Sprintf(MINIPROGRAM_GEN_ORDER_URL_FORMAT, client.PartnerCode, orderID)
	//需要增加query参数
	nSign := sign.NewSign(client.PartnerCode, client.CredentialCode)
	rawURL := nSign.GenSignURL(url)

	paramJSON, _ := json.Marshal(request)
	paramStr := string(paramJSON)

	_, _, errs := gorequest.New().Put(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		return false, MinipOrderResponse{}
	} else {
		return true, urlResp
	}
}

func (client *SailPayClient) GenNewMinipOrder(orderID string, request MinipOrderRequest) (gorequest.Response, string, []error) {

	url := fmt.Sprintf(MINIPROGRAM_GEN_ORDER_URL_FORMAT, client.PartnerCode, orderID)
	//需要增加query参数
	nSign := sign.NewSign(client.PartnerCode, client.CredentialCode)
	rawURL := nSign.GenSignURL(url)

	paramJSON, _ := json.Marshal(request)
	paramStr := string(paramJSON)

	return gorequest.New().Put(rawURL).Send(paramStr).End()
}
