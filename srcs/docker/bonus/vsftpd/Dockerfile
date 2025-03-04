FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				apk add vsftpd

COPY			docker/bonus/vsftpd/entry.sh /docker-entrypoint
COPY			docker/bonus/vsftpd/vsftpd.conf /etc/vsftpd/vsftpd.conf
RUN				chmod +x /docker-entrypoint

EXPOSE			21
EXPOSE			30000-30100

VOLUME			/var/ftp
STOPSIGNAL		SIGTERM

ENTRYPOINT		[ "/docker-entrypoint" ]

CMD				[ "vsftpd", "/etc/vsftpd/vsftpd.conf" ]
