FROM golang:1.15.0 as builder

WORKDIR /go/src/github.com/d-vignesh/shipper/vessel-service

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installprefix cgo main.go repository.go handler.go datastore.go

FROM debian:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/d-vignesh/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]
