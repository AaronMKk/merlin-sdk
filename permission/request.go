/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package permission
package permission

type RequestToCheckPermission struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
