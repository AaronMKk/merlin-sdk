package api

import (
	"fmt"
	"net/http"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/space"
)

const (
	baseSpaceUrl = "/v1/space"
)

func GetSpaceById(SpaceId string) (resp space.SpaceMetaDTO, code string, err error) {
	urlToGetSpaceById := fmt.Sprintf("%s/%s", baseSpaceUrl, SpaceId)

	req, err := http.NewRequest(http.MethodGet, httpclient.Endpoint(urlToGetSpaceById), nil)
	if err != nil {
		return
	}

	_, code, err = httpclient.Default.ForwardTo(req, &resp)

	return
}
