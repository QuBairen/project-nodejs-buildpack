package quote

import (
	"net/http"
)

// interface to allow unit testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
