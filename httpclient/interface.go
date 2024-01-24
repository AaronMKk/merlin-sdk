package httpclient

import (
	"fmt"
	"net/http"

	"encoding/json"
)

var (
	Default  HttpClient
	endpoint string
)

type HttpClient interface {
	ForwardTo(req *http.Request, jsonResp interface{}) (statusCode int, code string, err error)
}

func Init(endpointURL string) {
	Default = newhttpClient(3, 3)
	endpoint = endpointURL
}

func Endpoint(v string) string {
	return endpoint + v
}

type responseData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []byte `json:"data"`
}

func (resp *responseData) toResp(v interface{}) error {
	return json.Unmarshal(resp.Data, v)
}

func (resp *responseData) toError() error {
	if resp.Code == "" {
		return nil
	}

	return fmt.Errorf(resp.Msg)
}
