#!/bin/sh


echo ${REDIS_HOSTNAME}
if [ -f "/redis.ok" ]; then
	echo "[*] redis-cache already installed and configured"
else
	echo "[*] Installing redis-cache plugin"
	wp --allow-root config set WP_REDIS_HOST ${REDIS_HOSTNAME}
	wp --allow-root config set WP_REDIS_PORT ${REDIS_PORT}
	wp --allow-root config set WP_CACHE_KEY_SALT ${WP_URL}
	wp --allow-root plugin install redis-cache --activate
	wp --allow-root plugin update --all
	wp --allow-root redis enable
	touch /redis.ok
fi
