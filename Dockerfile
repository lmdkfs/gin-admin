FROM nicolaka/netshoot:latest
ADD ./bin/gin-admin /data/gin-admin

EXPOSE 8888

WORKDIR /data

