package sample_service

import (
	"context"

	desc "github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
)

type Service struct {
	desc.UnimplementedSampleServiceServer
}

func (i *Service) SampleMethod1(context.Context, *desc.SampleMethod1Request) (*desc.SampleMethod1Response, error) {
	return &desc.SampleMethod1Response{
		Value: &desc.Template{
			Id:   1,
			Name: "hello! it's a template",
		},
	}, nil
}
