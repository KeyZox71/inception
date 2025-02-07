# those are for production testing only do not execute in production environment

prod:
	docker compose -f $(DOCKER_CONTEXT)docker-compose.yml up -d --build

stop-prod:
	docker compose -f $(DOCKER_CONTEXT)docker-compose-prod.yml stop

.PHONY: prod stop-prod
