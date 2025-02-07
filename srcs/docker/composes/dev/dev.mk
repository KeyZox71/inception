# those rules a for developement purposes

dev:
	docker compose --profile dev -f $(DOCKER_CONTEXT)docker-compose.yml up -d --build

stop-dev:
	docker compose --profile dev -f $(DOCKER_CONTEXT)docker-compose.yml stop

.PHONY: dev stop-dev
