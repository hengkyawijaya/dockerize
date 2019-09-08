FROM golang:latest

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/env

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/env

RUN go build -o ./build/main  ./*.go

WORKDIR /build

ENV SERVICE_NAME=kamu

CMD ["./main"]

EXPOSE 8081