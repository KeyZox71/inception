DB_SERVER_NAME := mariadb
WEB_SERVER_NAME := nginx
CMS_NAME := wordpress
DOCKER_FOLDER := docker

DOCKER_CONTEXT := .
DOCKERFILE_DB := $(DOCKER_FOLDER)/$(DB_SERVER_NAME)/Dockerfile
DOCKERFILE_WEVSRV := $(DOCKER_FOLDER)/$(WEB_SERVER_NAME)/Dockerfile
DOCKERFILE_CMS := $(DOCKER_FOLDER)/$(CMS_NAME)/Dockerfile

mariadb-build:
	docker build -f $(DOCKERFILE_DB) -t $(DB_SERVER_NAME) .

nginx-build:
	docker build -f $(DOCKERFILE_WEVSRV) -t $(WEB_SERVER_NAME) .

wp-build:
	docker build -f $(DOCKERFILE_CMS) -t $(CMS_NAME) .
