FROM golang:1.15.0 as builder

WORKDIR /build

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM debian:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build/consignment-service .

CMD ["./consignment-service"]