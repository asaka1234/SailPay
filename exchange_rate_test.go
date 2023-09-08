package sailpay_client

import (
	"fmt"
	"testing"
)

func TestGetExchangeRate(t *testing.T) {

	client := NewSailPayClient(PartnerCode, CredentialCode)
	isSucceed, response := client.GetExchangeRate()
	fmt.Printf("%v %v", isSucceed, response)
}
