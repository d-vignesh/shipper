FROM golang:alphine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /build
# move working directory to /build
WORKDIR /build

ENV GO111MODULE=on

# copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy to code into the container
COPY . .

# build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-cli *.go

FROM alphine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /build/consignment-cli .
# COPY ./user-cli .

CMD ["./user-cli"]