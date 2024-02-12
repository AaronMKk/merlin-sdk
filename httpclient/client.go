package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

func newhttpClient(n int, timeout int) HttpClient {
	return &httpClient{
		MaxRetries: n,
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}},
		},
	}
}

// httpClient
type httpClient struct {
	Client     *http.Client
	MaxRetries int
}

func (hc *httpClient) ForwardTo(req *http.Request, jsonResp interface{}) (statusCode int, code string, err error) {
	req.Header.Set(config.TokenHeader, config.Token)

	resp, err := hc.do(req)
	if err != nil || resp == nil {
		return
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode

	v := new(responseData)
	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return
	}

	if err = v.toError(); err != nil {
		code = v.Code

		return
	}

	if jsonResp != nil {
		err = v.toResp(jsonResp)
	}

	return
}

func (hc *httpClient) do(req *http.Request) (resp *http.Response, err error) {
	if resp, err = hc.Client.Do(req); err == nil {
		return
	}

	maxRetries := hc.MaxRetries
	backoff := 10 * time.Millisecond

	for retries := 1; retries < maxRetries; retries++ {
		time.Sleep(backoff)
		backoff *= 2

		if resp, err = hc.Client.Do(req); err == nil {
			break
		}
	}
	return
}

func JsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(t); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
