#!/bin/sh

# Environment variables:
#   VSFTPD_USER: Username for the vsftpd user
#   VSFTPD_PASSWORD: Password for the vsftpd user

DONE_FILE=/vsftpd.ok

if [ -f "${DONE_FILE}" ]; then
	echo "[*] vsftpd already setup, skipping"
else
	echo "[*] Creating vsftpd user"

	adduser -D $VSFTPD_USER
    echo "$VSFTPD_USER:$VSFTPD_PASS" | /usr/sbin/chpasswd > /dev/null

	echo "[*] Giving vsftpd user ownership of WordPress data directory"
    chown -R "$VSFTPD_USER:$VSFTPD_USER" /var/ftp

	echo $VSFTPD_USER | tee -a /etc/vsftpd/vsftpd.userlist > /dev/null

	touch ${DONE_FILE}
fi

mkdir /var/run/vsftpd/empty -p

echo "[*] Starting '$@'"
exec "$@"
