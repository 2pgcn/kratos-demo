FROM registry.cn-shenzhen.aliyuncs.com/pg/debian:laster

ARG SERVER_FILE="bin/server"
WORKDIR /app

ADD  ${SERVER_FILE} /app/server

EXPOSE 8000
EXPOSE 9000
VOLUME /data/
RUN ["chmod","+x","/app/server"]
CMD ["./server", "-conf", "/data/conf"]
