#!/bin/bash

ps:
	docker compose ps

up:
	docker compose up -d

down:
	docker compose down

stop:
	docker compose stop

restart:
	docker compose restart

prune:
	docker system prune -f

build:
	docker compose build --no-cache

bash:
	docker compose exec -it app sh

fresh: down build up

force-fresh: down prune build up

test:
	docker compose exec -it app go test $(path)

help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  ps       - List the status of the Docker containers"
	@echo "  up       - Start the Docker containers in detached mode"
	@echo "  down     - Stop and remove the Docker containers"
	@echo "  stop     - Stop the Docker containers without removing them"
	@echo "  prune    - Remove unused Docker resources"
	@echo "  build    - Build the Docker images"
	@echo "  fresh    - Rebuild and restart the Docker containers"
	@echo "  test     - Run Go tests inside the app container (specify test path)"
	@echo "  help     - Display this help message"
