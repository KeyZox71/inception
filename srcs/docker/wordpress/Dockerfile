FROM			scratch AS builder
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				apk add go

WORKDIR			/build

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/wordpress/entrypoint

FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

COPY			--from=builder /build/entrypoint /docker-entrypoint
COPY			docker/wordpress/www-docker.conf /www-docker.conf

RUN				set -x \
				&& adduser -u 82 -D -S -G www-data www-data \
				&& apk add --no-cache curl tzdata fcgi \
				php84-phar php84-xml php84-curl php84-zip \
				php84-intl php84-mbstring php84-iconv \
				php84 php84-mysqli php84-fpm php84-json \
				php84-zlib php84-session php84-dom \
				php84-xmlreader php84-pdo php84-gd \
				php84-opcache php84-ctype php84-tokenizer \
				&& curl -O https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar \
				&& chmod +x wp-cli.phar \
				&& mv wp-cli.phar /usr/local/bin/wp \
				&& chmod +x /docker-entrypoint \
				&& rm -Rf /build \ 
				&& mkdir /docker-entrypoint.d

RUN				cp /usr/bin/php84 /usr/bin/php

VOLUME			/var/www/wordpress
VOLUME			/docker-entrypoint.d

WORKDIR			/var/www

RUN				version='6.7.2' \
				&& curl -o wordpress.tar.gz -fL "https://wordpress.org/wordpress-$version.tar.gz" \
				&& apk del curl \
				&& rm -rf /var/cache/apk/* \
				&& tar -xzvf wordpress.tar.gz \
				&& rm wordpress.tar.gz \
				&& chown -R www-data:www-data /var/www/wordpress \
				&& mkdir wp-content \
				&& for dir in /var/www/wp-content/*/ cache; do \
					dir="$(basename "${dir%/}")"; \
					mkdir "wp-content/$dir"; \
				done \
				&& chown -R www-data:www-data /var/www/wordpress/wp-content \
				&& chmod -R 1777 /var/www/wordpress/wp-content

ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/var/www/wordpress
STOPSIGNAL		SIGQUIT

EXPOSE			9000
CMD				[ "php-fpm84", "-F" ]
HEALTHCHECK		--interval=30s --timeout=10s --retries=3 --start-period=7s --start-interval=2s  CMD cgi-fcgi -bind -connect 127.0.0.1:${PHP_PORT:-9000} || exit 1
