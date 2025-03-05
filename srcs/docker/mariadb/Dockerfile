FROM			scratch AS builder
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				apk add go

WORKDIR			/build

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/mariadb/entrypoint \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/mariadb/healthcheck \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/getFileEnv

FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			version="0.1"
LABEL			maintainer="KeyZox"

RUN				addgroup mysql \
				&& adduser -S -G mysql mysql -h /var/lib/mysql

COPY			--from=builder /build/entrypoint /docker-entrypoint
COPY			--from=builder /build/healthcheck /docker-healthcheck
COPY			--from=builder /build/getFileEnv /bin/getFileEnv

RUN				apk add --no-cache mariadb tzdata mariadb-client \
				&& chmod +x /docker-healthcheck \
				&& chmod +x /docker-entrypoint \
				&& mkdir -p /etc/mysql/conf.d /etc/mysql/mariadb.conf.d/ /run/mariadb /run/mysqld \
				&& chmod ugo+rwx,o+t /run/mariadb \
				&& chown -R mysql:mysql /var/lib/mysql /run/mariadb /run/mysqld \
				&& rm -rf /var/cache/apk/*

ENV				LANG=C.UTF-8

COPY			--chmod=0644 configs/mariadb/docker.cnf /etc/my.cnf.d/mariadb-server.cnf.d

VOLUME			/var/lib/mysql

ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/var/lib/mysql

EXPOSE			3306
CMD				[ "mariadbd", "--user=mysql" ]
HEALTHCHECK		--interval=30s --timeout=10s --retries=3 --start-period=8s --start-interval=2s CMD /docker-healthcheck
