.PHONY: help setup migrate down push deploy

migrate:
	docker compose run --rm api sh -c 'migrate -source file:///app/internal/database/migrations -database "postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

setup:
	docker compose up --build -d --remove-orphans
	${MAKE} migrate

down:
	docker compose down --remove-orphans

push:
	docker build . -t teniolafatunmbi/go-todo
	docker login -u teniolafatunmbi
	docker push teniolafatunmbi/go-todo

deploy:
	helm install go-todo ./infra/helm/go-todo

help:
	@echo "setup - setup the environment"
	@echo "down - Tear down the dev environment"
	@echo "migrate - run database migrations"
	@echo "deploy - Deploy the app on Kubernetes with Helm"
