

This will be in two stages
1. Deployment with KinD for local testing
2. Deployment to AWS (use terraform to provision infrastructure)

## Deployment
- CD workflow that pushes an image to DockerHub before running `helm install`

### Steps
- Build the Go app
- Push the build to DockerHub
- Run `helm install` to deploy the application to k8s (on AWS or KinD)
