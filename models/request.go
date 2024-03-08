/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package models
package models

type ReqToResetLabel struct {
	Task       string   `yaml:"pipeline_tag"`
	Tags       []string `yaml:"tags"`
	Frameworks []string `yaml:"frameworks"`
}

// ModelDTO is a struct that represents a data transfer object for a model.
type ModelDTO struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Desc          string         `json:"desc"`
	Owner         string         `json:"owner"`
	Labels        ModelLabelsDTO `json:"labels"`
	Fullname      string         `json:"fullname"`
	CreatedAt     int64          `json:"created_at"`
	UpdatedAt     int64          `json:"updated_at"`
	LikeCount     int            `json:"like_count"`
	Visibility    string         `json:"visibility"`
	DownloadCount int            `json:"download_count"`
}

// ModelLabelsDTO is a struct that represents a data transfer object for model labels.
type ModelLabelsDTO struct {
	Task       string   `json:"task"`
	Others     []string `json:"others"`
	License    string   `json:"license"`
	Frameworks []string `json:"frameworks"`
}
