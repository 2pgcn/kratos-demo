FROM  registry.cn-shenzhen.aliyuncs.com/pg/golang:1.19.1 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn
RUN go build -gcflags="all=-N -l"  -o /src/server /src/cmd/...

FROM  registry.cn-shenzhen.aliyuncs.com/pg/debian:laster

WORKDIR /app

COPY --from=builder /src/server /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]