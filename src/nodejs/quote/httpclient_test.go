package quote_test

import (
	"net/http"
)

type MockClient struct {
	Dofunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}
