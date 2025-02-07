setup-corr:
	mkdir -P $HOME/data

correction:
	docker compose --profile correction -f $(DOCKER_CONTEXT)docker-compose.yml up -d --build

stop:
	docker compose --profile correction -f $(DOCKER_CONTEXT)docker-compose.yml stop

.PHONY: stop correction setup-corr
