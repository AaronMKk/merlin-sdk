package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/restfulauth"
)

var (
	restfulAuthUrl = "/v1/user/token/verify"
)

func CheckAuth(
	body restfulauth.RequestToRestfulAuth,
) (response restfulauth.ResponseToToRestfulAuth, code string, err error) {
	urlToRestfulAuth := httpclient.Endpoint(restfulAuthUrl)

	v, err := httpclient.JsonMarshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, urlToRestfulAuth, bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, response)

	return
}
