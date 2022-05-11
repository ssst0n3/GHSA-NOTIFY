FROM golang:1.17-alpine AS builder
# ENV GOPROXY="https://proxy.golang.org"
# ENV GOPROXY="https://goproxy.io,direct"
# COPY . /build
# WORKDIR /build
RUN GO111MODULE="on" GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" cmd/notifier
# RUN sed -i "s@http://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx notifier

FROM alpine:3.10
COPY --from=builder notifier /notifier
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /notifier
ENTRYPOINT ["/entrypoint.sh"]
