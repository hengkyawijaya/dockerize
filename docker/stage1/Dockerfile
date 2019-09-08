FROM golang:latest

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/docker

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/docker

RUN go build -o ./build/main  ./*.go

WORKDIR /build

ENV SERVICE_NAME=stage1

CMD ["./main"]

EXPOSE 8082