/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package restfulauth
package restfulauth

// RequestToRestfulAuth for CheckAuth request
type RequestToRestfulAuth struct {
	Token string `json:"token"`
}

// ResponseToToRestfulAuth for CheckAuth response
type ResponseToToRestfulAuth struct {
	Account string `json:"account"`
}
