.PHONY: gen
gen:
	buf generate protos

.PHONY: run-py
run-py:
	# https://pip.pypa.io/en/stable/installation/
    # python3 get-pip.py
    # pip install --upgrade google-api-python-client
	python3 ./python/write_proto.py

.PHONY: run-go
run-go:
	go run ./go/cmd/test