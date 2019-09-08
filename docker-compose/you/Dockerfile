FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/you

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/you

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure -v -vendor-only

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/main ./*.go

FROM scratch

COPY --from=builder go/src/github.com/hengkyawijaya/dockerize/you/build /build

WORKDIR /build

CMD ["./main"]

EXPOSE 8001
