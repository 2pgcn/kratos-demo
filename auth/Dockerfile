FROM golang:1.19.1 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn  go build -gcflags="all=-N -l"  -o /src/server /src/cmd/...

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /src/server /app
COPY --from=builder /src/configs/config.yaml /app/config.yaml

EXPOSE 8000
EXPOSE 9000

CMD ["./server", "-conf", "./config.yaml"]