GO111MODULE=off

.PHONY: run generate build dependencies

build:
	go build .

run:
	go run server/main.go

client:
	go run client/client.go

generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
