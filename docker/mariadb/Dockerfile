FROM			alpine:3.21

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				addgroup mysql \
				&& adduser -S -G mysql mysql -h /var/lib/mysql \
				&& echo 'permit nopass root as mysql' >> /etc/doas.conf

RUN				mkdir /build

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd

RUN				apk add --no-cache go \
				&& cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/mariadb-entry \
				&& cp /build/mariadb-entry /mariadb-entry \
				&& apk del go \
				&& rm -Rf /build

ENV				LANG=C.UTF-8

#COPY			--chmod=0644 docker.cnf /etc/my.cnf.d/

RUN				apk add --no-cache mariadb tzdata doas

VOLUME			/var/lib/mysql

#COPY			--chmod=0741 healthcheck.sh /healthcheck.sh

#ENTRYPOINT		[ "/mariadb-entry" ]
WORKDIR			/var/lib/mysql

USER			mysql
EXPOSE			3306
CMD				[ "mariadbd" ]
