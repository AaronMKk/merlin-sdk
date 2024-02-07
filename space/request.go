package space

type ReqToGetSpaceById struct {
	SpaceId string `json:"space_id"`
}

type SpaceMetaDTO struct {
	Id       string `json:"id"`
	SDK      string `json:"sdk"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Hardware string `json:"hardware"`
}
