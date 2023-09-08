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
查询订单详情
*/
func (client *SailPayClient) GetOrderInfo(request QueryOrderRequest) (bool, QueryOrderResponse) {

	var urlResp QueryOrderResponse

	url := QUERY_ORDER_URL

	//请求封装公共参数
	commonReq := CommonRequestInfo{
		MchNo:    client.MchNo,           //商户号
		AppId:    client.AppId,           //应用ID
		ReqTime:  time.Now().UnixMicro(), //请求时间
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
		QueryOrderRequest
	}
	result := UnifiedOrderRequestFinal{
		commonReq,
		request,
	}
	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)

	//发送请求
	_, _, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		return false, QueryOrderResponse{}
	} else {
		return true, urlResp
	}
}
