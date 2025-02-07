DB_SERVER_NAME := mariadb
WEB_SERVER_NAME := nginx
CMS_NAME := wordpress
DOCKER_FOLDER := docker

DOCKER_CONTEXT := srcs/
DOCKERFILE_DB := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(DB_SERVER_NAME)/Dockerfile
DOCKERFILE_WEVSRV := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(WEB_SERVER_NAME)/Dockerfile
DOCKERFILE_CMS := $(DOCKER_CONTEXT)$(DOCKER_FOLDER)/$(CMS_NAME)/Dockerfile

include srcs/docker/composes/dev/dev.mk
include srcs/docker/composes/correction.mk
include srcs/docker/composes/prod.mk

all: setup-corr correction

clean: stop
	docker system prune

fclean: clean
	docker system prune -a

re: clean all

.PHONY: all clean fclean re
