package sailpay_client

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"logtech.com/pay/SailPay/util/sign"
)

/*
获取当前汇率
*/
func (client *SailPayClient) GetExchangeRate() (bool, ExchangeRateResponse) {

	var urlResp ExchangeRateResponse

	url := fmt.Sprintf(MINIPROGRAM_QUERY_EXCHANGE_RATE, client.PartnerCode)
	//需要增加query参数
	nSign := sign.NewSign(client.PartnerCode, client.CredentialCode)
	rawURL := nSign.GenSignURL(url)

	_, _, errs := gorequest.New().Get(rawURL).EndStruct(&urlResp)
	if errs != nil {
		return false, ExchangeRateResponse{}
	} else {
		return true, urlResp
	}
}
