package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	api "github.com/ivch/testms/api/v1"
	httpClient "github.com/ivch/testms/client"
	"github.com/ivch/testms/client/parser"
	"github.com/ivch/testms/client/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	defaultGRPCHost      = "localhost"
	defaultGRPCPort      = "8081"
	defaultHTTPPort      = "8080"
	defaultClientTimeout = "15"
	defaultFilePath      = "../../ports.json"
)

func init() {
	if err := godotenv.Load("/run/secrets/client_env"); err != nil {
		if err := godotenv.Load(".env.dist"); err != nil {
			log.Fatal("unable to load env file")
		}
	}
}

func main() {
	grpcClient, timeout, ch, filepath, err := initDeps()
	if err != nil {
		log.Fatal("error creating dependecies: ", err)
	}

	err = parser.Start(filepath, ch)
	if err != nil {
		log.Fatal("error starting parser: ", err)
	}

	r := httpClient.New(service.New(grpcClient, timeout, ch))
	srv := &http.Server{
		Addr:         ":" + getenv("HTTP_PORT", defaultHTTPPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		log.Println("shutting down server...")
	}()

	fmt.Printf("started http server on :%s @ %s\n",
		getenv("HTTP_PORT", defaultHTTPPort), time.Now().Format("02/01/2006 15:04:05"))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func initDeps() (api.PortDomainServiceClient, time.Duration, chan api.SaveRequest, string, error) {
	t := getenv("CLIENT_TIMEOUT", defaultClientTimeout)
	intT, err := strconv.Atoi(t)
	if err != nil {
		return nil, 0, nil, "", err
	}

	timeout := time.Duration(int64(intT)) * time.Second

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s",
		getenv("GRPC_HOST", defaultGRPCHost),
		getenv("GRPC_PORT", defaultGRPCPort)),
		grpc.WithInsecure())
	if err != nil {
		return nil, 0, nil, "", err
	}
	grpcClient := api.NewPortDomainServiceClient(conn)

	//this should be set in config, i guess
	//in real system with limited RAM we should know approximate size of api.SaveRequest entity
	ch := make(chan api.SaveRequest, 10)

	fp := getenv("FILEPATH", defaultFilePath)

	return grpcClient, timeout, ch, fp, nil
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
