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

// 生成账单
/*
	orderID 商户内部订单id，要求同一商户唯一
*/
func (client *SailPayClient) PlaceUnifiedOrder(request UnifiedOrderRequest) (bool, UnifiedOrderResponse) {

	var urlResp UnifiedOrderResponse

	url := UNIFIEDORDER_URL

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
	fmt.Printf("sign = %+v\n", signVal)

	//合并复制
	type UnifiedOrderRequestFinal struct {
		CommonRequestInfo
		UnifiedOrderRequest
	}
	result := UnifiedOrderRequestFinal{
		commonReq,
		request,
	}

	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)
	fmt.Printf("json body=%+v+\n", paramStr)

	//发送请求
	_, _, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	if errs != nil {
		fmt.Printf("---1----%+v\n", errs)
		return false, UnifiedOrderResponse{}
	} else {
		fmt.Printf("---2----\n")
		return true, urlResp
	}
}
