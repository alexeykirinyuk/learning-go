package main

import (
	"log"
	"net"

	"github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/internal/app/sample_service"
	desc "github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	desc.RegisterSampleServiceServer(grpcServer, &sample_service.Service{})
	grpcServer.Serve(listener)
}
