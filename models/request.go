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
