DB_SERVER_NAME := mariadb
WEB_SERVER_NAME := nginx
CMS_NAME := wordpress
DOCKER_FOLDER := docker

DOCKER_CONTEXT := srcs/
DOCKERFILE_DB := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(DB_SERVER_NAME)/Dockerfile
DOCKERFILE_WEVSRV := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(WEB_SERVER_NAME)/Dockerfile
DOCKERFILE_CMS := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(CMS_NAME)/Dockerfile

dev:
	docker compose --profile dev -f $(DOCKER_CONTEXT)docker-compose.yml up -d --build

prod:
	docker compose -f $(DOCKER_CONTEXT)docker-compose.yml up -d --build

all: dev

stop:
	docker compose --profile dev -f $(DOCKER_CONTEXT)docker-compose.yml stop

stop-prod:
	docker compose -f $(DOCKER_CONTEXT)docker-compose-prod.yml stop

clean: stop
	docker system prune

fclean: clean
	docker system prune -a

re: clean all

.PHONY: cms-build db-build websrv-build clean-db clean-nginx
