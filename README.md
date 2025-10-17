## TODO API Deployment on Kubernetes

### Requirements
- Docker (v28.3.3 or newer)
- Kubernetes (v1.34 or newer)
- Helm (v3.18.4 or newer)

## Setup
- Pull this branch

- Run `make deploy`

- Run `kubectl get nodes -o wide` and copy the `INTERNAL-IP` of the node the `go-todo` namespace is on.

- Run `kubectl get svc` and copy the external port number for the `go-todo-svc` service.

- Run `curl <NODE-INTERNAL-IP>:<go-todo-svc-external-port>` to confirm that API is up. You should see
```sh
   { "message": "Hello World! Welcome to Go Todo" }
```

## Testing Endpoints
- Create todo

```sh

curl -X POST <NODE-INTERNAL-IP>:<go-todo-svc-external-port>/todos -H "Content-Type: application/json" -d '{"title": "Create test todo"}'

```

- Get todos

```sh

curl <NODE-INTERNAL-IP>:<go-todo-svc-external-port>/todos
```

- Update todo. You can play around with updating only the `title` or `is_completed` field.
```sh

 curl -X PUT <NODE-INTERNAL-IP>:<go-todo-svc-external-port>/todos/:id -H "Content-Type: application/json" -d '{"title": "Create test todo title update", "is_completed": true}'
```

- Delete todo

```sh

curl -X DELETE <NODE-INTERNAL-IP>:<go-todo-svc-external-port>:4000/todos/:id
```

# Notes
This Kubernetes deployment with Helm on KinD doesn't add an Ingress to the deployment. An Ingress will be added to the AWS deployment.
