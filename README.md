## BAREBONES TODO API WITH PostgreSQL

### Requirements
- Go (v 1.23.2+)
- Docker


## DB migration.
- Run `docker compose up -d`

- Run DB migration: 
```sh
migrate -source file://./internal/database/migrations -database "postgresql://postgres:postgres@localhost:5455/go_todo?sslmode=disable" up 2
```
- Run `go run main.go`

- Ping `localhost:4000` to confirm that API is up.

## TODO 
- Dockerise the API