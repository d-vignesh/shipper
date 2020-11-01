protos:
	protoc -I proto/ --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
	--go_out=proto/vessel --micro_out=proto/vessel	proto/vessel.proto

build:
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t vessel-service .

run:
	sudo docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns vessel-service