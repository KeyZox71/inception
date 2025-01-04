FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				groupadd --gid 999 -r mysql \
				&& useradd -r -g mysql mysql --home-dir /var/lib/mysql --uid 999 \
				&& echo 'permit nopass root as mysql' >> /etc/doas.conf

ENV				LANG C.UTF-8

COPY			--chmod=0644 docker.cnf /etc/my.cnf.d/

RUN				set -x \
				&& apk add --no-cache mariadb tzdata doas \

VOLUME			/var/lib/mysql

COPY			--chmod=0741 healthcheck.sh /healthcheck.sh
COPY			--chmod=0741 docker-entrypoint.sh /docker-entrypoint.sh

USER			mysql
EXPOSE			3306
WORKDIR			/var/lib/mysql
ENTRYPOINT		[ "/docker-entrypoint" ]
