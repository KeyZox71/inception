FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				set -x \
				&& adduser -u 82 -D -S -G www-data www-data \
				&& apk add --no-cache php84-fpm wget

VOLUME			/var/www/html

WORKDIR			/var/www/html
RUN				wget https://wordpress.org/wordpress-6.7.1.tar.gz \
				&& tar -xzvf wordpress-6.7.1.tar.gz \
				&& rm workpress-6.7.1.tar.gz

#COPY			docker-entrypoint.sh /
#RUN				[ "chmod", "+x", "/docker-entrypoint.sh" ]

#ENTRYPOINT		[ "/docker-entrypoint.sh" ]
WORKDIR			/var/www/html
STOPSIGNAL		SIGQUIT

EXPOSE			9000
CMD				[ "php-fpm84", "-F" ]
