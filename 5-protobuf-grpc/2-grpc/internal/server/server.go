package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/internal/app/sample_service"
	dsc "github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Cfg struct {
	Host     string
	GrpcPort string
	HttpPort string
}

type Server struct {
	cfg Cfg
}

func NewServer(cfg Cfg) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	grpcAddr := fmt.Sprintf("%s:%v", s.cfg.Host, s.cfg.GrpcPort)
	gatewayAddr := fmt.Sprintf("%s:%v", s.cfg.Host, s.cfg.HttpPort)

	gatewayServer := createGatewayServer(grpcAddr, gatewayAddr)

	go func() {
		log.Info().Msgf("Gateway server is running on %s", gatewayAddr)
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running gateway server")
		}
	}()

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("can't listen grpcAddr")
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			validateInterceptor,
		),
	}
	grpcServer := grpc.NewServer(opts...)

	dsc.RegisterSampleServiceServer(grpcServer, &sample_service.Service{})
	grpcServer.Serve(listener)
}
