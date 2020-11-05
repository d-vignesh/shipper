FROM golang:alphine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache 

RUN mkdir /build
WORKDIR /build

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service *.go

FROM alphine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build/vessel-service .
# COPY ./vessel-service .

CMD ["./vessel-service"]