package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/session"
)

var (
	urlToCheckAndRefresh = httpclient.Endpoint("/v1/session/check")
)

func CheckAndRefresh(param *session.RequestToCheckAndRefresh) (resp session.ResponseToCheckAndRefresh, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, urlToCheckAndRefresh, bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, _, err = httpclient.Default.ForwardTo(req, &resp)

	return
}