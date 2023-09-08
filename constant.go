package sailpay_client

type PayWayCode string
type CurrencyCode string
type DivisionMode int
type OrderState int
type PayDataType string

// 支付方式
const (
	SailCashier    PayWayCode = "SAIL_CASHIER"     //测试收单
	SailCashierBrl PayWayCode = "SAIL_CASHIER_BRL" //巴西收单
	SailCashierInr PayWayCode = "SAIL_CASHIER_INR" //印度收单
)

// 货币编码
const (
	Inr CurrencyCode = "inr" //印度卢比
)

// 分账模式
const (
	NoDivision     DivisionMode = 0 //- 该笔订单不允许分账[默认],
	AutoDivision   DivisionMode = 1 //- 支付成功按配置自动完成分账,
	ManualDivision DivisionMode = 2 //- 商户手动分账(解冻商户金额)
)

// 订单状态
const (
	StateCreate     OrderState = 0 //-订单生成
	StatePaying     OrderState = 1 //-支付中
	StatePaySucceed OrderState = 2 //-支付成功
	StatePayFailed  OrderState = 3 //-支付失败
	StateCanceled   OrderState = 4 //-已撤销
	StateChargeBack OrderState = 5 //-已退款
	StateClosed     OrderState = 6 //-订单关闭
)

// 货币编码
const (
	PayUrl     PayDataType = "payUrl"     //-跳转链接的方式
	Form       PayDataType = "form"       //-表单方式
	CodeUrl    PayDataType = "codeUrl"    //-二维码地址
	CodeImgUrl PayDataType = "codeImgUrl" //-二维码图片地址
	None       PayDataType = "none"       //-空支付参数
)

const (

	//1. 统一下单
	UNIFIEDORDER_URL = "https://pay.sailpayment.com/api/pay/unifiedOrder"
	//2. 查询订单
	QUERY_ORDER_URL = "https://query.sailpayment.com/api/pay/query"
	//查询余额
	QUERY_BALANCE_URL = "https://query.sailpayment.com/api/balance/query"
)

const (
	MchNo         = "test"                          //分配的商户号
	PrivateSecret = "EWEFD123RGSRETYDFNGFGFGSHDFGH" //分配的私钥
)
