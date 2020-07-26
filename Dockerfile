FROM golang:1.14-buster AS builder

ARG version

COPY ./ /go/prometheus-aio-filesd
WORKDIR /go/prometheus-aio-filesd
RUN go build -ldflags "-X main.version=$version" cmd/filesd.go -O prometheus-aio-filesd


FROM busybox:glibc

WORKDIR /app
COPY --from=builder /go/prometheus-aio-filesd/prometheus-aio-filesd /app

ENTRYPOINT ["/app/prometheus-aio-filesd"]