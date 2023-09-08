package sailpay_client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/exp/maps"
	"logtech.com/pay/SailPay/util/sign"
	"time"
)

/*
查询余额
*/
func (client *SailPayClient) GetBalanceInfo(request QueryBalanceRequest) (bool, QueryBalanceResponse) {

	var urlResp QueryBalanceResponse

	url := QUERY_BALANCE_URL

	//请求封装公共参数
	commonReq := CommonRequestInfo{
		MchNo:    client.MchNo,           //商户号
		AppId:    client.AppId,           //应用ID
		ReqTime:  time.Now().UnixMilli(), //请求时间
		Version:  "1.0",                  //接口版本号，固定：1.0
		SignType: "MD5",                  //签名类型，目前只支持MD5方式
	}

	//计算签名
	rawParams := structs.Map(request)
	commonParams := structs.Map(commonReq)
	maps.Copy(rawParams, commonParams)
	signVal := sign.GenSign(rawParams, client.PrivateSecret)
	commonReq.Sign = signVal //签名值
	fmt.Printf("rawSignStr = %+v\n", rawParams)
	fmt.Printf("sign = %+v\n", signVal)

	//合并复制
	type UnifiedOrderRequestFinal struct {
		CommonRequestInfo
		QueryBalanceRequest
	}
	result := UnifiedOrderRequestFinal{
		commonReq,
		request,
	}
	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)

	fmt.Printf("json body=%s\n", paramStr)

	//发送请求
	_, _, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		fmt.Printf("---1----%+v\n", errs)
		return false, QueryBalanceResponse{}
	} else {
		return true, urlResp
	}
}
