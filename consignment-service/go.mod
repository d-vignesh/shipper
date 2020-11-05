module github.com/d-vignesh/shipper/consignment-service

go 1.15

require (
	github.com/d-vignesh/shipper/vessel-service v0.0.0-20201101111948-42bfe90423bd
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
	google.golang.org/protobuf v1.25.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	<!-- github.com/d-vignesh/shipper/vessel-service => ../vessel-service -->
)
