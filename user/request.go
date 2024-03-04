/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package user
package user

type RequestToVerifyToken struct {
	Token string `json:"token"`
}

type ResponseToGetPlatFromUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
