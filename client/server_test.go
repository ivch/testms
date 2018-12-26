package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	api "github.com/ivch/testms/api/v1"
	client "github.com/ivch/testms/client/service"
	"google.golang.org/grpc"
)

type testGRPCClient struct{}

func (testGRPCClient) Save(ctx context.Context, in *api.SaveRequest, opts ...grpc.CallOption) (*api.SaveResponse, error) {
	return nil, nil
}

func (testGRPCClient) Get(ctx context.Context, in *api.GetRequest, opts ...grpc.CallOption) (*api.GetResponse, error) {
	if in.GetId() == "TESTT" {
		return &api.GetResponse{
			Name: "TEST1",
		}, nil
	}
	return nil, errors.New("some error to return")
}

func TestService_Get(t *testing.T) {
	h := New(initTestService())

	rr := httptest.NewRecorder()
	rq, _ := http.NewRequest("GET", "/v1/port/TESTT", nil)

	h.ServeHTTP(rr, rq)

	if rr.Code != http.StatusOK {
		t.Errorf("Wrong http response code: %d, instead of %d", rr.Code, http.StatusOK)
		return
	}

	rr = httptest.NewRecorder()
	rq, _ = http.NewRequest("GET", "/v1/port/TESST", nil)
	h.ServeHTTP(rr, rq)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Wrong http response code: %d, instead of %d", rr.Code, http.StatusInternalServerError)
		return
	}

	rr = httptest.NewRecorder()
	rq, _ = http.NewRequest("GET", "/v1/port/TEST1", nil)
	h.ServeHTTP(rr, rq)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Wrong http response code: %d, instead of %d", rr.Code, http.StatusNotFound)
		return
	}

	rr = httptest.NewRecorder()
	rq, _ = http.NewRequest("GET", "/v1/port/TEST", nil)
	h.ServeHTTP(rr, rq)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Wrong http response code: %d, instead of %d", rr.Code, http.StatusInternalServerError)
		return
	}

}

func initTestService() client.Service {
	ch := make(chan api.SaveRequest)
	return client.New(testGRPCClient{}, 15*time.Second, ch)
}
