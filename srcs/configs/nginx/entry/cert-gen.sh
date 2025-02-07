#!/bin/sh
if [ ! -f ${NGINX_SSL_KEY_FILE} ]; then
	echo "Generating certs"
	mkdir -p /etc/nginx/ssl
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${NGINX_SSL_KEY_FILE} -out ${NGINX_SSL_CERT_FILE} -subj "/C=FR/ST=IDF/L=Angouleme/O=42/OU=42/CN=adjoly.42.fr/UID=adjoly"
else 
	printf "Key already exist not recreating\n"
fi
