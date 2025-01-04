FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				set -x \
				&& addgroup -g 101 -S nginx\
				&& adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx \
				&& apk add --no-cache nginx tzdata \
				&& ln -sf /dev/stdout /var/log/nginx/access.log \
				&& ln -sf /dev/stderr /var/log/nginx/error.log

VOLUME			/etc/nginx
RUN				mkdir -p /etc/nginx/sites-available

COPY			docker-entrypoint.sh /docker-entrypoint.sh
RUN				[ "chmod", "+x", "/docker-entrypoint.sh" ]

ENTRYPOINT		[ "/docker-entrypoint.sh" ]
WORKDIR			/etc/nginx
STOPSIGNAL		SIGQUIT

EXPOSE			80
CMD				[ "nginx", "-g", "daemon off;" ]
