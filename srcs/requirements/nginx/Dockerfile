FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				set -x \
				&& apk update \
				&& addgroup -g 101 -S nginx\
				&& adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx \
				&& apk add nginx \
				&& ln -sf /dev/stdout /var/log/nginx/access.log \
				&& ln -sf /dev/stderr /var/log/nginx/error.log \
				&& apk add --no-cache tzdata

STOPSIGNAL		SIGQUIT

COPY			docker-entrypoint.sh /

ENTRYPOINT		[ "/docker-entrypoint.sh" ]

EXPOSE			80

CMD				[ "nginx", "-g", "daemon off;" ]
