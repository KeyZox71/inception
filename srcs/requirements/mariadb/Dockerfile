FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				groupadd -r mysql \
				&& useradd -r -g mysql mysql --home-dir /var/lib/mysql

RUN				set -x \
				&& 

VOLUME			/var/lib/mysql

COPY			healthcheck.sh /healthcheck.sh
COPY			docker-entrypoints.sh /docker-entrypoints.sh

ENTRYPOINT		["docker-entrypoint.sh"]

EXPOSE			3306
CMD				[ "mariadbd" ]
