package models

type ReqToResetLabel struct {
	Task       string   `yaml:"pipeline_tag"`
	Tags       []string `yaml:"tags"`
	Frameworks []string `yaml:"frameworks"`
}
