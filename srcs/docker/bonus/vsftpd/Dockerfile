FROM			scratch as builder
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				apk add --no-cache go

WORKDIR			/build

COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/vsftpd/entrypoint

FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

LABEL			maintainer="KeyZox"
LABEL			version="0.1"

COPY			--from=builder /build/entrypoint /docker-entrypoint
COPY			docker/bonus/vsftpd/vsftpd.conf /etc/vsftpd/vsftpd.conf

RUN				apk add vsftpd \
				&& mkdir -p /var/ftp \
				&& mkdir -p /docker-entrypoint.d

VOLUME			/var/ftp


ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/etc/vsftpd

EXPOSE			21
EXPOSE			30000-30100

STOPSIGNAL		SIGQUIT

CMD				[ "vsftpd", "/etc/vsftpd/vsftpd.conf" ]
