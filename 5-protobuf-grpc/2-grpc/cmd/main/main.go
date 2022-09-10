package main

import (
	"github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/internal/server"
)

func main() {
	s := server.NewServer(server.Cfg{
		Host:     "127.0.0.1",
		GrpcPort: "5002",
		HttpPort: "5005",
	})
	s.Run()
}
