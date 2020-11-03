protos:
	protoc -I proto/ --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
	--go_out=proto/user --micro_out=proto/user proto/user.proto 

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go
