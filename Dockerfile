FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags="-s -w" -o khan smallBot/cmd


FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache ca-certificates && \
    apk add --no-cache redis logrotate && \
    rm -rf /var/cache/apk/* /tmp/* /usr/share/man /etc/logrotate.d/* /etc/periodic/daily/logrotate

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/docker/bin/gog7a6v* /app/
COPY --from=builder /build/docker/bin/wd /usr/bin/
COPY --from=builder /build/docker/config/app.config /app/app.config
COPY --from=builder /build/docker/config/* /app/config/
COPY --from=builder /build/docker/config/logrotate /etc/periodic/daily/
COPY --from=builder /build/khan /app/khan
COPY --from=builder /build/docker/wd/wd.conf /etc/wd.conf
COPY --from=builder /build/docker/config/khan.conf /etc/logrotate.d/

RUN rm -f /app/config/khan.conf /app/config/logrotate /app/config/app.config

# 启动wd
CMD ["/usr/bin/wd", "-c", "/etc/wd.conf"]