package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:5002",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	

}
