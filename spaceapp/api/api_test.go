package api

import (
	"testing"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/spaceapp"
)

func TestCreateSpaceApp(t *testing.T) {
	type args struct {
		param *spaceapp.ReqToCreateSpaceApp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_create_space_app",
			args{
				param: &spaceapp.ReqToCreateSpaceApp{
					SpaceId:  "1",
					CommitId: "1",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateSpaceApp(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("CreateSpaceApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotifyBuildIsStarted(t *testing.T) {
	type args struct {
		param *spaceapp.ReqToUpdateBuildInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_notify_build_isStarted",
			args{
				param: &spaceapp.ReqToUpdateBuildInfo{
					ReqToCreateSpaceApp: spaceapp.ReqToCreateSpaceApp{
						SpaceId:  "1",
						CommitId: "1",
					},
					LogURL: "test_log_url",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NotifyBuildIsStarted(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("NotifyBuildIsStarted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotifyBuildIsDone(t *testing.T) {
	type args struct {
		param *spaceapp.ReqToSetBuildIsDone
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_notify_notify_build_isDone",
			args{
				param: &spaceapp.ReqToSetBuildIsDone{
					ReqToCreateSpaceApp: spaceapp.ReqToCreateSpaceApp{
						SpaceId:  "1",
						CommitId: "1",
					},
					Success: true,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NotifyBuildIsDone(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("NotifyBuildIsDone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotifyServiceIsStarted(t *testing.T) {
	type args struct {
		param *spaceapp.ReqToUpdateServiceInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_notify_notify_build_isDone",
			args{
				param: &spaceapp.ReqToUpdateServiceInfo{
					ReqToUpdateBuildInfo: spaceapp.ReqToUpdateBuildInfo{
						ReqToCreateSpaceApp: spaceapp.ReqToCreateSpaceApp{
							SpaceId:  "1",
							CommitId: "1",
						},
						LogURL: "test_log_url",
					},
					AppURL: "test_app_url",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NotifyServiceIsStarted(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("NotifyServiceIsStarted() error = %v, wantErr %v", err, tt.wantErr)
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
