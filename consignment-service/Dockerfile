# FROM golang:1.15.0 as builder

# ENV CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# WORKDIR /build

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY . . 

# RUN go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM debian:latest

RUN mkdir /app
WORKDIR /app

# COPY --from=builder /build/consignment-service .
COPY ./consignment-service .

CMD ["./consignment-service"]