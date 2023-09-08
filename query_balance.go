package sailpay_client

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"logtech.com/pay/SailPay/util/sign"
)

/*
查询余额
*/
func (client *SailPayClient) GetBalanceInfo(request QueryBalanceRequest) (bool, QueryBalanceResponse) {

	var urlResp QueryBalanceResponse

	url := QUERY_BALANCE_URL

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
		return false, QueryBalanceResponse{}
	} else {
		return true, urlResp
	}
}
