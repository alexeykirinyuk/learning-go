package main

import (
	"context"

	"github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:5002",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("error dial connect")
	}

	client := sample_service.NewSampleServiceClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"key", "stringValue",
		"key", "another-value",
		"another-key", "blah-blah",
	)

	resp, err := client.SampleMethod1(ctx, &sample_service.SampleMethod1Request{
		Id: 0,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("error response")
	}

	log.Debug().
		Interface("resp", resp).
		Msg("response received")
}
