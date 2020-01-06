FROM golang:1.13 AS builder

WORKDIR /go/src/github.com/ooclab/goproxy
COPY . .
RUN make build-static


FROM alpine:3.11.2
WORKDIR /work
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
COPY --from=builder /go/src/github.com/ooclab/goproxy/goproxy /usr/bin/goproxy
EXPOSE 8000
CMD ["/usr/bin/goproxy", "socks", "--listen=:8000"]
