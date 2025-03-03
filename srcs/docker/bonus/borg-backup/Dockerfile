FROM			scratch AS builder
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				apk add go

WORKDIR			/build

COPY			go.sum /build/go.sum
COPY			go.mod /build/go.mod
COPY			cmd /build/cmd
COPY			internal /build/internal

RUN				cd /build \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/borg-backup/entrypoint \
				&& go build git.keyzox.me/42_adjoly/inception/cmd/borg-backup/getpassphrase

FROM			scratch
ADD				docker/alpine/alpine-minirootfs-3.21.2-x86_64.tar.gz /

RUN				mkdir -p /backup \
				&& mkdir -p /source \
				&& mkdir /docker-entrypoint.d \
				&& mkdir /docker-backup.d

RUN				apk add --no-cache borgbackup tzdata \
				&& rm -rf /var/cache/apk/*

COPY			--from=builder /build/entrypoint /docker-entrypoint
COPY			--from=builder /build/getpassphrase /bin/getpassphrase
COPY			docker/bonus/borg-backup/default-bak.sh /docker-backup.d/default-bak.sh

ENTRYPOINT		[ "/docker-entrypoint" ]
WORKDIR			/

STOPSIGNAL		SIGQUIT

CMD				[ "crond", "-l", "8", "-f" ]
