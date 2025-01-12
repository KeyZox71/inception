DB_SERVER_NAME := mariadb
WEB_SERVER_NAME := nginx
CMS_NAME := wordpress
DOCKER_FOLDER := docker

DOCKER_CONTEXT := srcs/
DOCKERFILE_DB := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(DB_SERVER_NAME)/Dockerfile
DOCKERFILE_WEVSRV := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(WEB_SERVER_NAME)/Dockerfile
DOCKERFILE_CMS := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(CMS_NAME)/Dockerfile

build-db:
	docker build -f $(DOCKERFILE_DB) -t $(DB_SERVER_NAME) $(DOCKER_CONTEXT)

build-websrv:
	docker build -f $(DOCKERFILE_WEVSRV) -t $(WEB_SERVER_NAME) $(DOCKER_CONTEXT)

build-cms:
	docker build -f $(DOCKERFILE_CMS) -t $(CMS_NAME) $(DOCKER_CONTEXT)

start-db:
	docker compose -f $(DOCKER_CONTEXT)/docker-compose.yml up db --build 

start-nginx:
	docker compose -f $(DOCKER_CONTEXT)/docker-compose.yml up nginx --build 

start-wordp:
	docker compose -f $(DOCKER_CONTEXT)/docker-compose.yml up wordpress-php --build 

clean-db:
	docker stop inception-db
	docker container rm inception-db
	docker volume rm inception_wp-db
	docker image rm inception-db

clean-nginx:
	docker stop inception-nginx
	docker container rm inception-nginx

.PHONY: cms-build db-build websrv-build clean-db clean-nginx
