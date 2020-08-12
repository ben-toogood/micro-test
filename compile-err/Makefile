
GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/compile-err.proto
	
.PHONY: build
build:
	go build -o compile-err *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t compile-err:latest
