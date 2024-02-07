package api

import (
	"bytes"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/spaceapp"
)

var (
	urlToCreateSpaceApp         = "/v1/space-app"
	urlToNotifyBuildIsStarted   = "/v1/space-app/build/started"
	urlToNotifyBuildIsDone      = "/v1/space-app/build/done"
	urlToNotifyServiceIsStarted = "/v1/space-app/service/started"
)

func CreateSpaceApp(param *spaceapp.ReqToCreateSpaceApp) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, httpclient.Endpoint(urlToCreateSpaceApp), bytes.NewBuffer(v))
	if err != nil {
		return
	}
	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}

func NotifyBuildIsStarted(param *spaceapp.ReqToUpdateBuildInfo) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, httpclient.Endpoint(urlToNotifyBuildIsStarted), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}

func NotifyBuildIsDone(param *spaceapp.ReqToSetBuildIsDone) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, httpclient.Endpoint(urlToNotifyBuildIsDone), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}

func NotifyServiceIsStarted(param *spaceapp.ReqToUpdateServiceInfo) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, httpclient.Endpoint(urlToNotifyServiceIsStarted), bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}
