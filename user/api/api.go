package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/user"
)

const (
	urlToVerifyToken = "/v1/user/token/verify"
)

func VerifyToken(param *user.RequestToVerifyToken) (code string, err error) {
	if param == nil || param.Token == "" {
		code = "invalid param"
		err = fmt.Errorf("invalid param")
		return
	}

	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, httpclient.Endpoint(urlToVerifyToken), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}
