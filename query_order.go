package sailpay_client

import (
	"encoding/json"
	"fmt"
	"github.com/cyrilou/SailPay/util/sign"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/exp/maps"
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
		ReqTime:  time.Now().UnixMilli(), //请求时间
		Version:  "1.0",                  //接口版本号，固定：1.0
		SignType: "MD5",                  //签名类型，目前只支持MD5方式
	}

	//计算签名
	rawParams := structs.Map(request)
	commonParams := structs.Map(commonReq)
	maps.Copy(rawParams, commonParams)
	fmt.Printf("--1--%v\n", rawParams)
	signVal := sign.GenSign(rawParams, client.PrivateSecret)
	fmt.Printf("--2--\n")
	commonReq.Sign = signVal //签名值
	fmt.Printf("rawSignStr = %+v\n", rawParams)
	fmt.Printf("sign = %+v\n", signVal)

	//合并复制
	type QueryOrderRequestFinal struct {
		CommonRequestInfo
		QueryOrderRequest
	}
	result := QueryOrderRequestFinal{
		commonReq,
		request,
	}
	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)
	fmt.Printf("json body=%s\n", paramStr)

	//发送请求
	resp, body, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		fmt.Printf("---1----%+v\n", errs)
		fmt.Printf("---2----%+v\n", resp)
		fmt.Printf("---3----%+v\n", body)
		return false, QueryOrderResponse{}
	} else {
		fmt.Printf("---2----\n")
		return true, urlResp
	}
}
