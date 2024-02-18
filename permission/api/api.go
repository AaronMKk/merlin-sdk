package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/permission"
)

const (
	checkReadPermissionUrl   = "/v1/resource/permission/read"
	checkUpdatePermissionUrl = "/v1/resource/permission/update"
	checkDeletePermissionUrl = "/v1/resource/permission/delete"
)

func CanRead(
	body permission.RequestToCheckPermission,
) (code string, err error) {
	urlToCheckPermission := httpclient.Endpoint(checkReadPermissionUrl)

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

func CanUpdate(
	body permission.RequestToCheckPermission,
) (code string, err error) {
	urlToCheckPermission := httpclient.Endpoint(checkUpdatePermissionUrl)

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

func CanDelete(
	body permission.RequestToCheckPermission,
) (code string, err error) {
	urlToCheckPermission := httpclient.Endpoint(checkDeletePermissionUrl)

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
