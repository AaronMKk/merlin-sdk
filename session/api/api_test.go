package api

import (
	"reflect"
	"testing"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/session"
)

func TestCheckAndRefresh(t *testing.T) {
	type args struct {
		param *session.RequestToCheckAndRefresh
	}
	tests := []struct {
		name     string
		args     args
		wantResp session.ResponseToCheckAndRefresh
		wantErr  bool
	}{
		{
			"test_check_and_refresh",
			args{
				param: &session.RequestToCheckAndRefresh{
					IP:        "127.0.0.1",
					SessionId: "1",
					CSRFToken: "12345",
					UserAgent: "test1",
				},
			},
			session.ResponseToCheckAndRefresh{
				User:      "test1",
				CSRFToken: "12345",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, _, err := CheckAndRefresh(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAndRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CheckAndRefresh() gotResp = %v, want %v", gotResp, tt.wantResp)
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
