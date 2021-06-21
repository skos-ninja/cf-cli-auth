package cf

import "net/http"

type Transport interface {
	http.RoundTripper
}
