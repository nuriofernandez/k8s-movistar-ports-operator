> ⚠️ _Currently under development, expect breaking changes in the future._

# Movistar Ports K8S operator

This Kubernetes operator automates **Port Forwarding** and **NAT configuration** specifically for **Movistar HGU** routers. It allows you to manage external access to your home cluster using native K8s resources instead of manual router web-interface configuration.

Very useful if you are using [Movistar / O2](https://ipinfo.io/AS3352) and running K8S at home!

# How does it work?

This operator relies on [nuriofernandez/movistarapi](https://github.com/nuriofernandez/movistarapi) go library. Feel free to visit it if you will like to implement it on another project of yours!

Take into account both this operator and the movistarapi projects are under developement. Expect breaking changes!

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
  internalPort: 80
  host: 192.168.1.100
  protocol: TCP
```

After `kubectl apply -f <file>` your port will be open automatically!
