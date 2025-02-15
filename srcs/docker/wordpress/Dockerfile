FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				set -x \
				&& adduser -u 82 -D -S -G www-data www-data

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal
COPY			docker/wordpress/www-docker.conf /www-docker.conf
COPY			docker/wordpress/wp-config-docker.php /var/www/wordpress/wp-config-docker.php

RUN				apk add --no-cache go curl php84-mysqli php84-fpm tzdata fcgi \
				&& cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/wordpress/entrypoint \
				&& cp /build/entrypoint /docker-entrypoint \
				&& chmod +x /docker-entrypoint \
				&& apk del go \
				&& rm -Rf /build \ 
				&& mkdir -p /usr/src/wordpress

VOLUME			/var/www/wordpress

WORKDIR			/var/www

RUN				version='6.7.1' \
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
HEALTHCHECK		--interval=30s --timeout=10s --retries=3 --start-period=1s --start-interval=2s  CMD cgi-fcgi -bind -connect 127.0.0.1:${PHP_PORT:-9000} || exit 1
