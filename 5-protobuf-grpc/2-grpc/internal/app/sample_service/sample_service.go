package sample_service

import (
	"context"

	desc "github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
)

type Service struct {
	desc.UnimplementedSampleServiceServer
}

func (i *Service) SampleMethod1(ctx context.Context, req *desc.SampleMethod1Request) (*desc.SampleMethod1Response, error) {
	// return nil, status.Error(codes.InvalidArgument, "your request is bullshit")

	// md, ok := metadata.FromIncomingContext(ctx)
	// if ok {
	// 	log.Debug().
	// 		Interface("md", md).
	// 		Bool("ok", ok).
	// 		Msg("parsed metadata")
	// }

	return &desc.SampleMethod1Response{
		Value: &desc.Template{
			Id:   req.Id,
			Name: "hello! it's a template",
		},
	}, nil
}
