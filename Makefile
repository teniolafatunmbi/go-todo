.PHONY: setup migrate

migrate:
	docker compose run --rm api sh -c 'migrate -source file:///app/internal/database/migrations -database "postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

setup:
	docker compose up --build -d --remove-orphans
	${MAKE} migrate

down:
	docker compose down --remove-orphans

help:
	@echo "setup - setup the environment"
	@echo "down - Tear down the dev environment"
	@echo "migrate - run database migrations"

