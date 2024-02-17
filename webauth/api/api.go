package client

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/webauth"
)

var (
	webAuthUrl = "/v1/session/check"
)

func CheckAuth(
	body webauth.RequestToAuth,
) (resp webauth.ResponseToAuth, code string, err error) {
	urlToWebAuth := httpclient.Endpoint(webAuthUrl)

	v, err := httpclient.JsonMarshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, urlToWebAuth, bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, resp)

	return
}
