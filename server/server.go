package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	api "github.com/ivch/testms/api/v1"
	"google.golang.org/grpc"
)

//Run starts grpc request listener
func Run(ctx context.Context, v1API api.PortDomainServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterPortDomainServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("starting gRPC server on :%s @ %s\n", port, time.Now().Format("02/01/2006 15:04:05"))
	return server.Serve(listen)
}
