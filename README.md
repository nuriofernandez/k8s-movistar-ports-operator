> ⚠️ _Currently under development, expect breaking changes in the future._

# Movistar Ports K8S operator

This kubernetes operator handles the port opening for Movistar routers.

Useful if you are using K8S at home!

# How to install?

For first time run:
```bash
git clone github.com/nuriofernandez/k8s-movistar-ports-operator
cd k8s-movistar-ports-operator
kubectl apply -f config/crd/movistarport.yaml
```

Moving forward you need to keep the operator running:
```bash
MOVISTAR_ROUTER_PASS=%replace-with-router-pass% go run ./cmd/main.go
```
> Note: Right now, there is no container for this, you need to manually run this until I continue with the development.

# Example resource:

```yaml
apiVersion: nurio.me/v1alpha1
kind: MovistarPort
metadata:
  name: HttpPort
spec:
  externalPort: 80
  protocol: TCP
```

After `kubectl apply -f <file>` your port will be open automatically!
