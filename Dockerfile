FROM registry.cn-shenzhen.aliyuncs.com/pg/debian:laster

COPY ${data} /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
