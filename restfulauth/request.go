package restfulauth

type RequestToRestfulAuth struct {
	Token string `json:"token"`
}

type ResponseToToRestfulAuth struct {
	Account string `json:"account"`
}
