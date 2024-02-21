package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/permission"
)

const (
	checkReadPermissionUrl   = "/v1/coderepo/permission/read"
	checkUpdatePermissionUrl = "/v1/coderepo/permission/update"
	checkDeletePermissionUrl = "/v1/coderepo/permission/delete"
)

// CanRead for request resource permission read
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

// CanUpdate for request resource permission update
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

// CanDelete for request resource permission delete
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
