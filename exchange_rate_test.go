package globepay_client

import (
	"fmt"
	"testing"
)

func TestGetExchangeRate(t *testing.T) {

	client := NewGlobePayClient(PartnerCode, CredentialCode)
	isSucceed, response := client.GetExchangeRate()
	fmt.Printf("%v %v", isSucceed, response)
}
