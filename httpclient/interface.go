package httpclient

import (
	"fmt"
	"net/http"

	"encoding/json"
)

const (
	maxRetries = 3
	timeout    = 10
)

var (
	config  Config
	Default HttpClient
)

type HttpClient interface {
	ForwardTo(req *http.Request, jsonResp interface{}) (statusCode int, code string, err error)
}

func Init(cfg *Config) {
	Default = newhttpClient(maxRetries, timeout)
	config = *cfg
}

// Endpoint
func Endpoint(v string) string {
	return config.Endpoint + v
}

// Config
type Config struct {
	Token       string `json:"token"          required:"true"`
	Endpoint    string `json:"endpoint"       required:"true"`
	TokenHeader string `json:"token_header"   required:"true"`
}

// responseData
type responseData struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
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
