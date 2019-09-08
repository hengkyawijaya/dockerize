FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/stage2

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/stage2

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/main ./*.go

FROM scratch

COPY --from=builder go/src/github.com/hengkyawijaya/dockerize/docker/stage2/build /build

WORKDIR /build

ENV SERVICE_NAME=stage2

CMD ["./main"]

EXPOSE 8083
