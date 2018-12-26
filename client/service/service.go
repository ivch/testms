package service

import (
	"context"
	"log"
	"time"

	api "github.com/ivch/testms/api/v1"
)

//Service describes client interface
type Service interface {
	Get(name string) (*api.GetResponse, error)
}

type service struct {
	portsClient   api.PortDomainServiceClient
	clientTimeout time.Duration
	data          <-chan api.SaveRequest
}

//New creates service instance and runs process() function which reads data from the given channel
//and sends it to the server
func New(client api.PortDomainServiceClient, timeout time.Duration, ch <-chan api.SaveRequest) Service {
	s := service{
		portsClient:   client,
		clientTimeout: timeout,
		data:          ch,
	}

	go s.process()

	return &s
}

func (s *service) Get(id string) (*api.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.clientTimeout)
	defer cancel()

	req := api.GetRequest{
		Id: id,
	}

	res, err := s.portsClient.Get(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) save(req api.SaveRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.portsClient.Save(ctx, &req)
	if err != nil {
		//here i've put only stdlib logging but error handling should be business specific
		log.Println("error saving data: ", err)
	}
}

//process incoming data channel
//regarding to requirements with 200mb RAM i suggested that device would be limited in CPU too
//so maybe there is no reason of making this processing in goroutines
func (s *service) process() {
	for req := range s.data {
		s.save(req)
	}
}

//nevertheless i've added this variant of process function, maybe i'm wrong about the device
//if devices are different - waitgroup size should be passed to New func
//func (s *service) process() {
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//
//			for req := range s.data {
//				s.save(req)
//			}
//		}()
//	}
//
//	wg.Wait()
//}
