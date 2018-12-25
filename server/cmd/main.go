package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ivch/testms/server"
	"github.com/ivch/testms/server/service"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to load env file")
	}
}

func main() {
	if os.Getenv("GRPC_PORT") == "" {
		log.Fatal("wrong port given")
	}

	ctx := context.Background()

	db, err := initDeps()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	svc := service.New(db)

	if err := server.Run(ctx, svc, os.Getenv("GRPC_PORT")); err != nil {
		log.Fatal(err)
	}
}

func initDeps() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SCHEMA")))
	if err != nil {
		return nil, err
	}

	return db, nil
}
