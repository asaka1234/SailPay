package globepay_client

type GlobePayClient struct {
	PartnerCode    string
	CredentialCode string
}

// 实例化请求端
func GlobePayClient(partnerCode string, credentialCode string) *GlobePayClient {
	var client GlobePayClient
	client.PartnerCode = partnerCode
	client.CredentialCode = credentialCode
	return &client
}
