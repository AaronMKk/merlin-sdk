package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/models"
)

const (
	baseModelUrl = "/v1/model"
)

func ResetLabel(modelId string, param *models.ReqToResetLabel) (code string, err error) {
	urlToResetLabel := httpclient.Endpoint(fmt.Sprintf("%s/%s/label", baseModelUrl, modelId))

	v, err := httpclient.JsonMarshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, urlToResetLabel, bytes.NewBuffer(v))
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, nil)

	return
}

// GetModelById get model info by id
func GetModelById(modelId string) (resp models.ModelDTO, code string, err error) {
	urlToGetById := httpclient.Endpoint(fmt.Sprintf("%s/%s", baseModelUrl, modelId))

	req, err := http.NewRequest(http.MethodGet, urlToGetById, nil)
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, &resp)

	return
}
