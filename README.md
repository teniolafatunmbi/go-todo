## BAREBONES TODO API WITH EPHEMERAL DATA STORE

## DB migration.
- Run DB migration with 
```sh
migrate -source file://./internal/database/migrations -database "postgresql://postgres:postgres@localhost:5455/go_todo?sslmode=disable" up 2
```
