package api

import (
	"reflect"
	"testing"

	"github.com/openmerlin/merlin-sdk/httpclient"
	"github.com/openmerlin/merlin-sdk/space"
)

func TestGetSpaceById(t *testing.T) {
	type args struct {
		SpaceId string
	}
	tests := []struct {
		name     string
		args     args
		wantResp space.SpaceMetaDTO
		wantErr  bool
	}{
		{
			"test_helloT123",
			args{
				SpaceId: "1",
			},
			space.SpaceMetaDTO{
				Id:       "1",
				SDK:      "docker",
				Name:     "helloT123",
				Owner:    "test1",
				Hardware: "cpu basic 2 vcpu · 16gb · free",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, _, err := GetSpaceById(tt.args.SpaceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpaceById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetSpaceById() gotResp = %v, want %v", gotResp, tt.wantResp)
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
