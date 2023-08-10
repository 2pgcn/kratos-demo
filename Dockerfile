FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

ARG SERVER_FILE="bin/server"
WORKDIR /app

ADD  ${SERVER_FILE} /app/server

EXPOSE 8000
EXPOSE 9000
VOLUME /data/

RUN ["chmod","+x","/app/server"]
CMD ["./server", "-conf", "/data/conf"]
