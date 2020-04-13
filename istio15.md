```
docker.io/istio/operator:1.5.1
docker.io/istio/kubectl:1.5.1
docker.io/istio/pilot:1.5.1
```

```
oc new-project cloudos-mesh
oc adm policy add-scc-to-group anyuid system:serviceaccounts:cloudos-mesh
helm template install/kubernetes/helm/istio-init --name istio-init --namespace cloudos-mesh | kubectl apply -f -


```
