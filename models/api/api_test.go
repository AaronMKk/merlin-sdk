package api

import (
	"testing"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/models"
)

func TestResetLabel(t *testing.T) {
	type args struct {
		modelId string
		param   *models.ReqToResetLabel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_reset_label",
			args{
				modelId: "1",
				param: &models.ReqToResetLabel{
					Task:       "test task",
					Tags:       []string{"test tag"},
					Frameworks: []string{"test frameworks"},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ResetLabel(tt.args.modelId, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("ResetLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// 在测试之前执行一些初始化操作
	cf := httpclient.Config{
		Token:       "12345",
		Endpoint:    "http://localhost:8888/internal",
		TokenHeader: "Token",
	}
	httpclient.Init(&cf)

	// 运行测试套件
	m.Run()
}
