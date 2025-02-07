#!/bin/sh
if [ ! -v ${PRODUCTION} ]; then
	if [ ! -f ${NGINX_SSL_KEY_FILE} ]; then
		echo "Generating certs"
		mkdir -p /etc/nginx/ssl
		openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${NGINX_SSL_KEY_FILE} -out ${NGINX_SSL_CERT_FILE} -subj "/C=FR/ST=IDF/L=Angouleme/O=42/OU=42/CN=adjoly.42.fr/UID=adjoly"
	else 
		printf "Key already exist not recreating\n"
	fi
else
	printf "Entering production mode for nginx"
	INPUT_FILE="/etc/nginx/http.d/www.conf"
	OUTPUT_FILE="/etc/nginx/http.d/www.conf"
	sed -E '
	s/listen\s+443 ssl;/listen 80;/; 
	s/server_name.*/&\n\tlisten 80;/; 
	/ssl_certificate/d; 
	/ssl_certificate_key/d; 
	/ssl_protocols/d; 
	/ssl_session_timeout/d;
	' "$INPUT_FILE" > "$OUTPUT_FILE"
fi
