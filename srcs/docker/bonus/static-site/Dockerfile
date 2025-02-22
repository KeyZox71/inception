FROM			inception-nginx

LABEL			maintainer="KeyZox"


COPY			kanelovh/ /var/www/html/
COPY			--chmod=644 www.conf.template /etc/nginx/templates/www.conf.template
COPY			--chmod=644 cert.sh /docker-entrypoint.d/
