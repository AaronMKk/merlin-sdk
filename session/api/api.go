package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/session"
)

const (
	urlToClear           = "/v1/session/clear"
	urlToCheckAndRefresh = "/v1/session/check"
)

func CheckAndRefresh(param *session.RequestToCheckAndRefresh) (resp session.ResponseToCheckAndRefresh,
	code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, httpclient.Endpoint(urlToCheckAndRefresh), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, &resp)

	return
}

func Clear(param *session.RequestToClear) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodDelete, httpclient.Endpoint(urlToClear), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}
