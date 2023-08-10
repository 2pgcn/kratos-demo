FROM  registry.cn-shenzhen.aliyuncs.com/pg/golang:1.20 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM registry.cn-shenzhen.aliyuncs.com/pg/debian:laster

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
