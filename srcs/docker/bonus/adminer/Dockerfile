FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

VOLUME			/var/www/adminer

COPY			docker/bonus/adminer/adminer-4.17.1-mysql.php /var/www/adminer/index.php

RUN				apk update \
				&& apk add php84 php84-mysqli php84-session curl \
				&& addgroup -S adminer \
				&& adduser -S -G adminer adminer \
				&& chown -R adminer:adminer /var/www/adminer \
				&& chmod 755 /var/www/adminer/index.php

RUN				echo "upload_max_filesize = 128M" >> /etc/php84/conf.d/0-upload_large_dumps.ini \
				&&	echo "post_max_size = 128M" >> /etc/php84/conf.d/0-upload_large_dumps.ini \
				&&	echo "memory_limit = 1G" >> /etc/php84/conf.d/0-upload_large_dumps.ini \
				&&	echo "max_execution_time = 600" >> /etc/php84/conf.d/0-upload_large_dumps.ini \
				&&	echo "max_input_vars = 5000" >> /etc/php84/conf.d/0-upload_large_dumps.ini

WORKDIR			"/var/www/adminer"

EXPOSE			8080
STOPSIGNAL		SIGINT

USER			adminer
CMD				[ "php84", "-S", "0.0.0.0:8080", "-t", "/var/www/adminer"]
HEALTHCHECK		--interval=30s --timeout=10s --retries=3 --start-period=1s --start-interval=2s CMD curl -f "http://localhost:8080"
