package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/user"
)

const (
	urlToVerifyToken  = "/v1/user/token/verify"
	urlToPlatformInfo = "/v1/user/%s/platform"
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

func GetPlatFromUser(account string) (resp, code string, err error) {
	url := fmt.Sprintf(urlToPlatformInfo, account)
	req, err := http.NewRequest(http.MethodGet, httpclient.Endpoint(url), nil)
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, &resp)

	return
}
