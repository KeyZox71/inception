FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				groupadd --gid 999 -r mysql \
				&& useradd -r -g mysql mysql --home-dir /var/lib/mysql --uid 999

ENV				LANG C.UTF-8

RUN				set -x \
				&& apk add --no-cache mariadb gosu \
				&& gosu --version \
				&& gosu nobody true

VOLUME			/var/lib/mysql

COPY			healthcheck.sh /healthcheck.sh
COPY			docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT		[ "docker-entrypoint.sh" ]

EXPOSE			3306
CMD				[ "mariadbd" ]
