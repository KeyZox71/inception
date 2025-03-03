FROM			scratch AS builder
ADD				alpine-minirootfs-3.21.2-x86_64.tar.gz /

WORKDIR			/build

COPY			go.sum /build/go.sum
COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/borg-backup/entrypoint

FROM			scratch
ADD				alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				mkdir -p /backup \
				&& mkdir -p /source \
				&& mkdir /docker-entrypoint.d \
				&& mkdir /docker-backup.d

RUN				apk add --no-cache borgbackup tzdata \
				&& rm -rf /var/cache/apk/*

COPY			--from=builder /build/entrypoint /docker-entrypoint
COPY			docker/bonus/borg-backup/default-bak.sh /docker-backup.d

ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/

STOPSIGNAL		SIGQUIT

CMD				[ "crond", "-l", "${CRON_LOGLEVEL:-8}", "-f" ]
