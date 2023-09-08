package sailpay_client

type SailPayClient struct {
	MchNo         string //商户号
	PrivateSecret string //私钥
	AppId         string //商户号
}

// 实例化请求端
func NewSailPayClient(mchNo string, appId string, privateSecret string) *SailPayClient {
	var client SailPayClient
	client.MchNo = mchNo
	client.AppId = appId
	client.PrivateSecret = privateSecret
	return &client
}

// 公共参数
type CommonRequestInfo struct {
	MchNo    string `json:"mchNo" structs:"mchNo"`       //商户号
	AppId    string `json:"appId" structs:"appId"`       //应用ID
	ReqTime  int64  `json:"reqTime" structs:"reqTime"`   //请求接口时间,13位时间戳
	Version  string `json:"version" structs:"version"`   //接口版本号，固定：1.0
	SignType string `json:"signType" structs:"signType"` //签名类型，目前只支持MD5方式
	//签名值
	Sign string `json:"sign" structs:"sign"` //签名值

}

// 统一订单-请求
type UnifiedOrderRequest struct {
	MchOrderNo string `json:"mchOrderNo" structs:"mchOrderNo"` //商户生成的订单号
	WayCode    string `json:"wayCode" structs:"wayCode"`       //支付方式
	Amount     int    `json:"amount" structs:"amount"`         //支付金额,单位分
	Currency   string `json:"currency" structs:"currency"`     //三位货币代码,印度卢比:inr
	Subject    string `json:"subject" structs:"subject"`       //商品标题
	Body       string `json:"body" structs:"body"`             //商品描述
	NotifyUrl  string `json:"notifyUrl" structs:"notifyUrl"`   //支付结果异步回调URL
	ReturnUrl  string `json:"returnUrl" structs:"returnUrl"`   //支付结果同步跳转通知URL
	//Version     string       `json:"version"`     //接口版本号，固定：1.0
	//SignType    string       `json:"signType"`    //签名类型，目前只支持MD5方式
	ExtParam string `json:"extParam" structs:"extParam"` //商户扩展参数json格式字符串 至少有country参数字段 ,回调时会原样返回

	//可选参数
	ExpiredTime  int    `json:"expiredTime,omitempty" structs:"expiredTime,omitempty"`   //订单失效时间,单位秒,默认2小时.订单在(创建时间+失效时间)后失效
	ClientIp     string `json:"clientIp,omitempty" structs:"clientIp,omitempty"`         //客户端IPV4地址
	ChannelExtra string `json:"channelExtra,omitempty" structs:"channelExtra,omitempty"` //特定渠道发起的额外参数,json格式字符串.详见渠道参数说明
	DivisionMode int    `json:"divisionMode,omitempty" structs:"divisionMode,omitempty"` //分账模式
}

// 统一订单-返回
type UnifiedOrderResponse struct {
	Code int `json:"code"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在
	//可选字段
	Msg  string           `json:"msg"`  //商户编码
	Sign string           `json:"sign"` //对data内数据签名,如data为空则不返回
	Data UnifiedOrderData `json:"data"` //返回下单数据,json格式数据
}

// 统一订单-data
type UnifiedOrderData struct {
	//正确的返回
	PayOrderId  string `json:"payOrderId"`  //支付订单号
	MchOrderNo  string `json:"mchOrderNo"`  //商户传入的订单号
	OrderState  int    `json:"orderState"`  //支付订单状态
	PayDataType string `json:"payDataType"` //支付参数类型
	//可选返回
	PayData string `json:"payData"` //发起支付用到的支付参数
	ErrCode string `json:"errCode"` //上游渠道返回的错误码
	ErrMsg  string `json:"errMsg"`  //上游渠道返回的错误描述
}

//-------------------------------------------------------

// 查询订单
type QueryOrderRequest struct {
	PayOrderId string `json:"payOrderId" structs:"payOrderId"` //支付中心生成的订单号，与mchOrderNo二者传一即可
	MchOrderNo string `json:"mchOrderNo" structs:"mchOrderNo"` //商户生成的订单号，与payOrderId二者传一即可
}

// 查询订单-返回
type QueryOrderResponse struct {
	Code int `json:"code"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在
	//可选字段
	Msg  string         `json:"msg"`  //商户编码
	Sign string         `json:"sign"` //对data内数据签名,如data为空则不返回
	Data QueryOrderData `json:"data"` //返回下单数据,json格式数据
}

// 查询订单-data
type QueryOrderData struct {
	//正确的返回
	PayOrderId string `json:"payOrderId"` //返回支付系统订单号
	MchNo      string `json:"mchNo"`      //商户号
	AppId      string `json:"appId"`      //应用ID
	MchOrderNo string `json:"mchOrderNo"` //返回商户传入的订单号
	IfCode     string `json:"ifCode"`     //支付接口编码
	WayCode    string `json:"wayCode"`    //支付方式,如支付链接URL SAIL_CASHIER
	Amount     int    `json:"amount"`     //支付金额,单位分
	Currency   string `json:"currency"`   //三位货币代码,印度:inr
	State      int    `json:"state"`      //支付订单状态
	Subject    string `json:"subject"`    //商品标题
	Body       string `json:"body"`       //商品描述
	CreatedAt  int64  `json:"createdAt"`  //订单创建时间,13位时间戳
	//可选返回
	ClientIp       string `json:"clientIp"`       //客户端IPV4地址
	ChannelOrderNo string `json:"channelOrderNo"` //对应渠道的订单号
	ErrCode        string `json:"errCode"`        //渠道下单返回错误码
	ErrMsg         string `json:"errMsg"`         //渠道下单返回错误描述
	ExtParam       string `json:"extParam"`       //商户扩展参数,回调时会原样返回
	SuccessTime    int64  `json:"successTime"`    //订单支付成功时间,13位时间戳
}

//-------------------------------------------------------

// 查询余额
type QueryBalanceRequest struct {
	Currency string `json:"currency" structs:"currency"` //三位货币编号,如 巴西雷亚尔 BRL
}

// 查询余额-返回
type QueryBalanceResponse struct {
	Code int `json:"code"` //SUCCESS表示创建订单成功，EXISTS表示订单已存在
	//可选字段
	Msg  string           `json:"msg"`  //商户编码
	Sign string           `json:"sign"` //对data内数据签名,如data为空则不返回
	Data QueryBalanceData `json:"data"` //返回下单数据,json格式数据
}

// 查询余额-data
type QueryBalanceData struct {
	//正确的返回
	MchNo       string `json:"mchNo"`       //商户号
	AppId       string `json:"appId"`       //应用ID
	Currency    string `json:"currency"`    //三位货币代码,巴西雷亚尔:BRL
	Balance     string `json:"balance"`     //总余额 单位分
	Available   string `json:"available"`   //可用余额 单位分
	Unavailable string `json:"unavailable"` //不可用余额 单位分
	Frozen      string `json:"frozen"`      //冻结资金 单位分
}

//------------回调返回的参数说明--------------------------

// 回调收到的数据
type UnifiedOrderNotify struct {
	PayOrderId string `json:"payOrderId" structs:"payOrderId"` //商户生成的订单号
	MchNo      string `json:"mchNo" structs:"mchNo"`           //商户号
	AppId      string `json:"appId" structs:"appId"`           //应用ID
	MchOrderNo string `json:"mchOrderNo" structs:"mchOrderNo"` //商户生成的订单号
	IfCode     string `json:"ifCode" structs:"ifCode"`         //支付接口编码
	WayCode    string `json:"wayCode" structs:"wayCode"`       //支付方式
	Amount     int    `json:"amount" structs:"amount"`         //支付金额,单位分
	Currency   string `json:"currency" structs:"currency"`     //三位货币代码,印度卢比:inr
	State      int    `json:"state" structs:"state"`           //支付订单状态
	Subject    string `json:"subject" structs:"subject"`       //商品标题
	Body       string `json:"body" structs:"body"`             //商品描述
	ExtParam   string `json:"extParam" structs:"extParam"`     //商户扩展参数,回调时会原样返回
	CreatedAt  int64  `json:"createdAt" structs:"createdAt"`   //订单创建时间,13位时间戳
	ReqTime    string `json:"reqTime" structs:"reqTime"`       //请求接口时间,13位时间戳
	//可选返回
	ClientIp       string `json:"clientIp,omitempty" structs:"clientIp,omitempty"`       //客户端IPV4地址
	ChannelOrderNo string `json:"channelOrderNo" structs:"channelOrderNo"`               //对应渠道的订单号
	ErrCode        string `json:"errCode,omitempty" structs:"errCode,omitempty"`         //渠道下单返回错误码
	ErrMsg         string `json:"errMsg,omitempty" structs:"errMsg,omitempty"`           //渠道下单返回错误描述
	SuccessTime    int64  `json:"successTime,omitempty" structs:"successTime,omitempty"` //订单支付成功时间,13位时间戳

	//签名值
	Sign string `json:"sign" structs:"sign"` //签名值
}

//-------------------------------------------------------
