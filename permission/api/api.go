package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/permission"
)

var (
	checkPermissionUrl = "/v1/permission/check"
)

func CheckPermission(
	body permission.RequestToCheckPermission,
) (code string, err error) {
	urlToCheckPermission := httpclient.Endpoint(checkPermissionUrl)

	v, err := httpclient.JsonMarshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, urlToCheckPermission, bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return

}
