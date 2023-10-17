package sailpay_client

import (
	"encoding/json"
	"fmt"
	"github.com/cyrildou/SailPay/util/logger"
	"github.com/cyrildou/SailPay/util/sign"
	"github.com/fatih/structs"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/exp/maps"
	"time"
)

/*
查询订单结算列表或者未结算列表
*/
func (client *SailPayClient) QuerySettleList(request QuerySettleListRequest) (bool, QuerySettleListResponse) {

	var urlResp QuerySettleListResponse

	url := QUERY_SETTLE_LIST_URL

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
	//fmt.Printf("rawSignStr = %+v\n", rawParams)
	//fmt.Printf("sign = %+v\n", signVal)

	//合并复制
	type QuerySettleListRequestFinal struct {
		CommonRequestInfo
		QuerySettleListRequest
	}
	result := QuerySettleListRequestFinal{
		commonReq,
		request,
	}
	//构造请求body
	paramJSON, _ := json.Marshal(result)
	paramStr := string(paramJSON)
	//fmt.Printf("json body=%s\n", paramStr)

	//发送请求
	resp, body, errs := gorequest.New().Post(url).Send(paramStr).EndStruct(&urlResp)
	logger.Infof("SailPaySdk-settle_list: [reqUrl]%s [reqBody]%s  [respBody]%s, [resp]%+v, [err]:%+v\n", url, paramStr, body, resp, errs)
	if errs != nil {
		fmt.Printf("err: body:%s, resp:%+v, info:%+v\n", body, resp, errs)
		return false, QuerySettleListResponse{}
	} else {
		return true, urlResp
	}
}
