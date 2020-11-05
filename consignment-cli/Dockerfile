FROM golang:alphine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

# move working directory to /build
RUN mkdir /build
WORKDIR /build

ENV GO111MODULE=on

# copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy to code into the container
COPY . .

# build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-cli .

FROM alphine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /build/consignment-cli .
COPY --from=builder /build/consignment.json .
# COPY ./consignment-cli .
# COPY ./consignment.json .

CMD ["./consignment-cli"]