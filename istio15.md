```
docker.io/istio/operator:1.5.1

docker.io/istio/kubectl:1.5.1
docker.io/istio/pilot:1.5.1
Image:         docker.io/istio/pilot:1.5.1
Image:         docker.io/istio/proxyv2:1.5.1
Image:         docker.io/jaegertracing/all-in-one:1.16
Image:         docker.io/prom/prometheus:v2.15.1
Image:         grafana/grafana:6.5.2
Image:         quay.io/kiali/kiali:v1.15

```

```
oc new-project cloudos-mesh
oc adm policy add-scc-to-group anyuid system:serviceaccounts:cloudos-mesh
helm template install/kubernetes/helm/istio-init --name istio-init --namespace cloudos-mesh | kubectl apply -f -


```
