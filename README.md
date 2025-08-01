## BAREBONES TODO API WITH PostgreSQL + Containerization

### Requirements
- Docker

## Setup
- Run `cp .env.example .env`

- Run `make setup`

- Ping `localhost:4000` to confirm that API is up. You should see 
```sh
   { "message": "Hello World! Welcome to Go Todo" }
```

## Testing Endpoints
- Create todo

```sh

curl -X POST localhost:4000/todos -H "Content-Type: application/json" -d '{"title": "Create test todo"}'

```

- Get todos

```sh

curl localhost:4000/todos
```

- Update todo. You can play around with updating only the `title` or `is_completed` field.
```sh

 curl -X PUT localhost:4000/todos/:id -H "Content-Type: application/json" -d '{"title": "Create test todo title update", "is_completed": true}'
```

- Delete todo

```sh

curl -X DELETE localhost:4000/todos/:id 
```
