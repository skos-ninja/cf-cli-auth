package cf

import (
	"errors"
	"net/http"
)

const clientIDHeader = "CF-Access-Client-Id"
const clientSecretHeader = "CF-Access-Client-Secret"

type ServiceTokenClient struct {
	clientId, clientSecret string
	tr                     http.RoundTripper
}

func (c *ServiceTokenClient) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.clientId == "" {
		return nil, errors.New("missing client id")
	}
	if c.clientSecret == "" {
		return nil, errors.New("missing client secret")
	}

	req.Header.Set(clientIDHeader, c.clientId)
	req.Header.Set(clientSecretHeader, c.clientSecret)

	resp, err := c.tr.RoundTrip(req)
	return resp, err
}

func NewServiceTokenClient(tr http.RoundTripper, clientId, clientSecret string) Transport {
	return &ServiceTokenClient{
		clientId:     clientId,
		clientSecret: clientSecret,
		tr:           tr,
	}
}
