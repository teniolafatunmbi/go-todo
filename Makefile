.PHONY: setup migrate

migrate:
	docker compose run --rm api sh -c 'migrate -source file:///app/internal/database/migrations -database "postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up 2'

setup:
	docker compose up -d --remove-orphans
	${MAKE} migrate

help:
	@echo "setup - setup the environment"
	@echo "migrate - run database migrations"
