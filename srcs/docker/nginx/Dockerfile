FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				set -x \
				&& addgroup -g 101 -S nginx\
				&& adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx \
				&& apk add --no-cache go nginx openssl tzdata envsubst curl \
				&& cd /build \
				&& go get git.keyzox.me/42_adjoly/inception/cmd/nginx/entrypoint \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/nginx/entrypoint \
				&& cp /build/entrypoint /docker-entrypoint \
				&& chmod +x /docker-entrypoint \
				&& apk del go \
				&& ln -sf /dev/stdout /var/log/nginx/access.log \
				&& ln -sf /dev/stderr /var/log/nginx/error.log \
				&& rm -Rf /build \ 
				&& rm -rf /var/cache/apk/* \
				&& mkdir -p /var/www/html \
				&& mkdir -p /etc/nginx/conf.d \
				&& mkdir /docker-entrypoint.d

VOLUME			/etc/nginx/templates
VOLUME			/docker-entrypoint.d
COPY			docker/nginx/health.conf /etc/nginx/http.d/health.conf

ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/etc/nginx

STOPSIGNAL		SIGQUIT

EXPOSE			80
CMD				[ "nginx", "-g", "daemon off;" ]
HEALTHCHECK		--interval=30s --timeout=10s --retries=3 --start-period=3s --start-interval=2s CMD curl http://127.0.0.1:80/health || exit 1
