

This will be in two stages
1. Deployment with KinD for local testing
2. Deployment to AWS (use terraform to provision infrastructure)

## Deployment
- CD workflow that pushes an image to DockerHub before running `helm install`

### Todos
- Build the Go app ✅
- Push the build to DockerHub✅
- Configure k8s deployment with Helm (on KinD):
  - Setup DB migration✅
  - Ensure that DB migration runs before main app deployment (use wait-for)
  - Setup main app deployment with DB✅
- Configure k8s deployment with Helm (on AWS)
- Write documentation
