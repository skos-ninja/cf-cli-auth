package access

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func Get(ctx context.Context, appAud, appDomain string) (string, error) {
	// We try to return the token from disk before performing an auth session
	t, err := GetLocalToken(ctx, appAud, appDomain)
	if err != nil {
		return "", err
	} else if t != "" {
		return t, nil
	}

	return "", errors.New("not implemented")
}

func IsAccessResponse(resp *http.Response) bool {
	if resp == nil || resp.StatusCode != http.StatusFound {
		return false
	}

	location, err := resp.Location()
	if err != nil || location == nil {
		return false
	}
	if strings.HasPrefix(location.Path, "/cdn-cgi/access/login") {
		return true
	}

	return false
}
