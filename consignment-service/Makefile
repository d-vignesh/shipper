.PHONY: protos

protos:
	protoc -I proto/ --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
	--go_out=proto/consignment --micro_out=proto/consignment proto/consignment.proto 

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	# sudo docker build -t consignment-service .

run :
	sudo docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service