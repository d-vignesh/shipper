# FROM golang:1.15.0 as builder

# # set necessary environment variable needed for our image
# ENV CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# # move working directory to /build
# WORKDIR /build

# # copy and download dependency using go mod
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # copy to code into the container
# COPY . .

# # build the application
# RUN go build -a -installsuffix cgo -o consignment-cli .

FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

# COPY --from=builder /build/consignment-cli .
COPY ./consignment-cli .
COPY ./consignment.json .

CMD ["./consignment-cli"]