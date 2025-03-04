FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				apk update \
				&& apk add redis

RUN				echo "maxmemory 256mb" >> /etc/redis.conf \
				&& echo "maxmemory-policy allkeys-lru" >> /etc/redis.conf \
				&& sed -i -r "s/bind 127.0.0.1//" /etc/redis.conf

EXPOSE			6379
CMD				[ "redis-server", "--protected-mode", "no"]
