module github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/3-grpc-client

go 1.19

require (
	github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.28.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/alexeykirinyuk/learning-go/5-protobuf-grpc/2-grpc/pkg/sample_service => ./../2-grpc/pkg/sample_service
