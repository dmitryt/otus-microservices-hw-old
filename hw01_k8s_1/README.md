# Working with Kubernetes

## Task

1) Create minimalistic web service, which will run on port 8000. It should have endpoint GET /health and should respond with `{"status": "OK"}`
2) Build locally docker image and push it to Dockerhub
3) Create k8s manifests for following entities: Deployment, Service, Ingress. Deployment can contain live and rediness probes. The amount of replicas is >= 2.
Ingress host should be `arch.homework`. Afterall GET `http://arch.homework/health` should respond with `{"status": "OK"}`

Additional task: Ingress should contain the rule, which forward all requests `/otusapp/{student name}/*` to service with path rewrites

### Apply k8s manifests
```bash
kubectl apply -f k8s
```
### Testing
1. Example of sending request to cluster(Ingress IP address can be different)
```bash
curl -H 'Host: arch.homework' http://192.168.64.2/health

{"status":"OK"}
```
2. Request to any unsupported route returns 404 from Ingress
```bash
curl -H 'Host: arch.homework' http://192.168.64.2/another-route

<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx</center>
</body>
</html>
```
3. Request to /otusapp/dmitry.tsilynko/ should redirect to service. All other requests with the same prefix will be handled by internal app
```bash
curl -H 'Host: arch.homework' http://192.168.64.2/otusapp/dmitry.tsilynko/health

{"status":"OK"}

curl -H 'Host: arch.homework' http://192.168.64.2/otusapp/dmitry.tsilynko/another-route

404 page not found
```
