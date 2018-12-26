package service

import (
	"context"
	"errors"
	"testing"
	"time"

	api "github.com/ivch/testms/api/v1"
	"google.golang.org/grpc"
)

type testGRPCClient struct{}

func (testGRPCClient) Save(ctx context.Context, in *api.SaveRequest, opts ...grpc.CallOption) (*api.SaveResponse, error) {
	if in.GetId() != "TESTT" {
		return nil, errors.New("some error to return")
	}

	return &api.SaveResponse{}, nil
}

func (testGRPCClient) Get(ctx context.Context, in *api.GetRequest, opts ...grpc.CallOption) (*api.GetResponse, error) {
	if in.GetId() != "TESTT" {
		return nil, errors.New("some error to return")
	}

	return &api.GetResponse{}, nil

}

func TestService_Get(t *testing.T) {
	srv := initTestService()

	if _, err := srv.Get("TESTT"); err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if _, err := srv.Get("TEST1"); err == nil {
		t.Error("expected error got nil")
	}
}

func initTestService() Service {
	ch := make(chan api.SaveRequest)
	return New(testGRPCClient{}, 15*time.Second, ch)
}
