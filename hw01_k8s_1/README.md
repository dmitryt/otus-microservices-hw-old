# Working with Kubernetes

## Task

1) Create minimalistic web service, which will run on port 8000. It should have endpoint GET /health and should respond with `{"status": "OK"}`
2) Build locally docker image and push it to Dockerhub
3) Create k8s manifests for following entities: Deployment, Service, Ingress. Deployment can contain live and rediness probes. The amount of replicas is >= 2.
Ingress host should be `arch.homework`. Afterall GET `http://arch.homework/health` should respond with `{"status": "OK"}`

Additional task: Ingress should contain the rule, which forward all requests `/otusapp/{student name}/*` to service with path rewrites
