
migrate:
	docker compose run -it api sh -c "migrate -source file://./internal/database/migrations -database postgres://$DB_USER:$DB_PASS&$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"