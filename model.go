package globepay_client

type GlobePayClient struct {
	PartnerCode    string
	CredentialCode string
}

// 实例化请求端
func NewGlobePayClient(partnerCode string, credentialCode string) *GlobePayClient {
	var client GlobePayClient
	client.PartnerCode = partnerCode
	client.CredentialCode = credentialCode
	return &client
}

//小程序订单-请求
type MinipOrderRequest struct {
	Description string `json:"description"` //订单标题（最大长度128字符，超出自动截取）
	Price       int    `json:"price"`       //金额，单位为货币最小单位，例如使用100表示GBP1.00
	Currency    string `json:"currency"`    //币种代码 默认值: GBP,允许值: GBP, CNY
	NotifyURL   string `json:"notify_url"`  //支付通知url，详见支付通知api，不填则不会推送支付通知
	Operator    string `json:"operator"`    //操作人员标识
	Appid       string `json:"appid"`       //小程序appid
	CustomerID  string `json:"customer_id"` //小程序openid
}

//小程序订单-返回
/**
 * 返回信息:
 * return_code          return_msg
 * --------------------------------------
 * ORDER_NOT_EXIST      订单不存在
 * ORDER_MISMATCH       订单号与商户不匹配
 * SYSTEMERROR          系统内部异常
 * INVALID_SHORT_ID     商户编码不合法或没有对应商户
 * SIGN_TIMEOUT         签名超时，time字段与服务器时间相差超过5分钟
 * INVALID_SIGN         签名错误
 * PARAM_INVALID        参数不符合要求，具体细节可参考return_msg字段
 * NOT_PERMITTED        未开通网关支付权限
 * --------------------------------------
 */
type MinipOrderResponse struct {
	ReturnCode string `json:"return_code"` //执行结果

	//正确的返回
	ResultCode     string `json:"result_code"`      //SUCCESS表示创建订单成功，EXISTS表示订单已存在
	PartnerCode    string `json:"partner_code"`     //商户编码
	FullName       string `json:"full_name"`        //商户注册全名
	PartnerName    string `json:"partner_name"`     //商户名称
	OrderID        string `json:"order_id"`         //GlobePay订单ID，同时也是微信订单ID，最终支付成功的订单ID可能不同
	PartnerOrderID string `json:"partner_order_id"` //商户订单ID
	SDKParams      string `json:"sdk_params"`       //小程序openid

	//错误的返回
	ReturnMsg string `json:"return_msg"`
}

//汇率
type ExchangeRateResponse struct {
	ReturnCode string `json:"return_code"` //执行结果

	//正确的返回
	WechatRate       float64 `json:"wechat_rate"`        //微信汇率
	AlipayRetailRate float64 `json:"alipay_retail_rate"` //支付宝线下渠道汇率
	AlipayOnlineRate float64 `json:"alipay_online_rate"` //支付宝线上渠道汇率

	//错误的返回
	ReturnMsg string `json:"return_msg"`
}
