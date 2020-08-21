FROM nicolaka/netshoot:latest
ADD ./bin/gin-admin /data/gin-admin
RUN apk add --no-cache stress-ng bcc-tools bcc-doc  # /usr/share/bcc
EXPOSE 8888

WORKDIR /data

