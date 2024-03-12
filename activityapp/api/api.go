package api

import (
	"bytes"
	"github.com/openmerlin/merlin-sdk/activityapp"
	"github.com/openmerlin/merlin-sdk/httpclient"
	"net/http"
)

const (
	urlToModifyActivity = "/v1/activity"
)

func CreateActivity(param *activityapp.ReqToCreateActivity) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, httpclient.Endpoint(urlToModifyActivity), bytes.NewBuffer(v))
	if err != nil {
		return
	}
	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}

func DeleteActivity(param *activityapp.ReqToCreateActivity) (code string, err error) {
	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, httpclient.Endpoint(urlToModifyActivity), bytes.NewBuffer(v))
	if err != nil {
		return
	}
	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}
