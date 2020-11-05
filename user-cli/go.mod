module github.com/d-vignesh/shipper/user-cli

go 1.15

require (
	github.com/d-vignesh/shipper/user-service v0.0.0-20201103145152-c1b54512d698
	github.com/micro/cli v0.2.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
)

replace (
	<!-- github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible -->
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
