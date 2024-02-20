/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package spaceapp
package spaceapp

type ReqToCreateSpaceApp struct {
	SpaceId  string `json:"space_id"`
	CommitId string `json:"commit_id"`
}

type ReqToUpdateBuildInfo struct {
	ReqToCreateSpaceApp

	LogURL string `json:"log_url"`
}

type ReqToSetBuildIsDone struct {
	ReqToCreateSpaceApp

	Success bool `json:"success"`
}

type ReqToUpdateServiceInfo struct {
	ReqToUpdateBuildInfo

	AppURL string `json:"app_url"`
}
