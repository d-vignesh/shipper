FROM golang:1.15.0 as builder

WORKDIR /go/src/github.com/d-vignesh/shipper/consignment-cli

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder

CMD ["./consignment-cli"]