package http

import (
	"github.com/MG-RAST/Shock/vendor/github.com/MG-RAST/golib/stretchr/testify/mock"
	"net/http"
)

type TestRoundTripper struct {
	mock.Mock
}

func (t *TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := t.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
