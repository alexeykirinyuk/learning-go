.PHONY: gen
gen:
	buf generate api

.PHONY: run
run:
	go run ./cmd/main

.PHONE: ui
ui:
	grpcui --proto ./api/sample_service/sample_service.proto --plaintext localhost:5002