package cf

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflared/token"
	"github.com/rs/zerolog"
)

const jwtTokenHeader = "Cf-Access-Token"

type AccessTokenClient struct {
	appDomain, token string
	tr               http.RoundTripper
}

func NewAccessTokenClient(ctx context.Context, tr http.RoundTripper, appDomain string) (Transport, error) {
	u, err := url.Parse(appDomain)
	if err != nil {
		return nil, err
	}

	appInfo, err := token.GetAppInfo(u)
	if err != nil {
		return nil, err
	}

	t, err := token.FetchTokenWithRedirect(u, appInfo, zerolog.Ctx(ctx))
	if err != nil {
		return nil, err
	}

	return &AccessTokenClient{
		appDomain: u.Hostname(),
		token:     t,
		tr:        tr,
	}, nil
}

func (c *AccessTokenClient) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Hostname() == c.appDomain {
		if c.token == "" {
			return nil, errors.New("access token missing")
		}

		req.Header.Set(jwtTokenHeader, c.token)
	} else {
		log.Printf("Hit url not required: %s\n", req.URL.String())
	}

	resp, err := c.tr.RoundTrip(req)
	return resp, err
}
