#!/bin/sh

WP_DIR="/var/www/wordpress"
#WP_VERSION="6.7.2"

#if [ -f "$WP_DIR/index.php" ]; then
#	echo "Already downloaded wordpress, skipping..."
#else
#	echo "Downloading Wordpress ${WP_VERSION}"
#	wp --allow-root core download --version=${WP_VERSION} --path=${WP_DIR}
#fi


if [ -f "${WP_DIR}/wp-config.php" ]; then
	echo "Wordpress already configured, skipping installation"
else
	echo "Confiruring wordpress..."
	sleep 1
	echo $TZ
	wp --allow-root core config --dbname=${WP_DB_NAME} --dbuser=${WP_DB_USER} --dbpass=${WP_DB_PASS} --dbhost=${WP_DB_HOST} --dbprefix=wp_ --path=${WP_DIR}
	wp --allow-root core install --url=https://${WP_URL} --title="${WP_TITLE}" --admin_user=${WP_ADMIN_USER} --admin_password=${WP_ADMIN_PASS} --admin_email=${WP_ADMIN_EMAIL} --path=${WP_DIR}
	wp option update blog_public ${WP_SEARCH_ENGINE_VISIBILITY} --allow-root
fi
