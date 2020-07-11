```
http://www.talkwithtrend.com/Article/246491
```
```
istio x-envoy-decorator-operation wrong
```

```
[root@cloudos03 ~]# curl -v 10.244.1.125:32099
* About to connect() to 10.244.1.125 port 32099 (#0)
*   Trying 10.244.1.125...
* Connected to 10.244.1.125 (10.244.1.125) port 32099 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 10.244.1.125:32099
> Accept: */*
> 
< HTTP/1.1 503 Service Unavailable
< content-length: 33
< content-type: text/plain
< date: Fri, 08 May 2020 11:43:36 GMT
< server: envoy
< x-envoy-decorator-operation: library-userinfo-istio-svc.spaceubeewz96.svc.cluster.local:32099/*
< 
* Connection #0 to host 10.244.1.125 left intact

root@cloudos03 istio-component]# curl -v  10.100.187.153:32101
* About to connect() to 10.100.187.153 port 32101 (#0)
*   Trying 10.100.187.153...
* Connected to 10.100.187.153 (10.100.187.153) port 32101 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 10.100.187.153:32101
> Accept: */*
> 
< HTTP/1.1 404 
< Content-Type: application/json;charset=UTF-8
< Transfer-Encoding: chunked
< Date: Fri, 08 May 2020 12:23:48 GMT
< 
* Connection #0 to host 10.100.187.153 left intact


[root@cloudos03 istio-component]# nslookup istio-policy.cloudos-mesh.svc.cloudos
Server:		10.99.20.3
Address:	10.99.20.3#53

Name:	istio-policy.cloudos-mesh.svc.cloudos
Address: 10.100.149.177

[root@cloudos03 istio-component]# nslookup library-userinfo-istio-svc.spaceubeewz96.svc.cloudos.local
Server:		10.99.20.3
Address:	10.99.20.3#53

** server can't find library-userinfo-istio-svc.spaceubeewz96.svc.cloudos.local: REFUSED




[root@cloudos03 ~]# kubectl exec -it library-userinfo-istio-svc-85dc9c6bf6-vt52q -n spaceubeewz96  -c istio-proxy /bin/bash
istio-proxy@library-userinfo-istio-svc-85dc9c6bf6-vt52q:/$ curl 127.0.0.1:15000/clusters
curl: (7) Failed to connect to 127.0.0.1 port 15000: Connection refused
istio-proxy@library-userinfo-istio-svc-85dc9c6bf6-vt52q:/$ curl http://127.0.0.1:15000/config_dump
curl: (7) Failed to connect to 127.0.0.1 port 15000: Connection refused



"bootstrap": {
   "@type": "type.googleapis.com/envoy.admin.v2alpha.BootstrapConfigDump",
   "bootstrap": {
    "node": {
     "id": "sidecar~10.244.1.125~library-userinfo-istio-svc-85dc9c6bf6-vt52q.spaceubeewz96~spaceubeewz96.svc.cluster.local",
     "cluster": "library-userinfo-istio",
     "metadata": {
      "openshift.io/scc": "privileged",
      "pod-template-hash": "4187572692",
      "POD_NAME": "library-userinfo-istio-svc-85dc9c6bf6-vt52q",
      "INTERCEPTION_MODE": "REDIRECT",
      "istio": "sidecar",
      "ISTIO_PROXY_VERSION": "1.0.2",
      "ISTIO_PROXY_SHA": "istio-proxy:930841ca88b15365737acb7eddeea6733d4f98b9",
      "ISTIO_VERSION": "1.0.4",
      "release": "userinfo-spaceubeewz96",
      "app": "library-userinfo-istio"
     },
     "build_version": "0/1.8.0-dev//RELEASE"
    },
  


```
hpa
```
NAMESPACE      NAME                   REFERENCE                         TARGETS         MINPODS   MAXPODS   REPLICAS   AGE
cloudos-mesh   istio-egressgateway    Deployment/istio-egressgateway    <unknown>/80%   1         5         1          3h
cloudos-mesh   istio-ingressgateway   Deployment/istio-ingressgateway   <unknown>/80%   1         5         1          3h
cloudos-mesh   istio-pilot            Deployment/istio-pilot            <unknown>/80%   1         5         1          3h
cloudos-mesh   istio-policy           Deployment/istio-policy           <unknown>/80%   1         5         1          3h
cloudos-mesh   istio-telemetry        Deployment/istio-telemetry        <unknown>/80%   1         5         1          3h

```
```

https://github.com/istio/istio/issues/10335
+ ip6tables -A INPUT -j REJECT
ip6tables v1.6.0: can't initialize ip6tables table `filter': Table does not exist (do you need to insmod?)
Perhaps ip6tables or your kernel needs to be upgraded.

--allow-privileged=true to k8s apiserver  


+ ip6tables -F INPUT
ip6tables v1.6.0: can't initialize ip6tables table `filter': Table does not exist (do you need to insmod?)
Perhaps ip6tables or your kernel needs to be upgraded.
+ true
+ ip6tables -A INPUT -m state --state ESTABLISHED -j ACCEPT
ip6tables v1.6.0: can't initialize ip6tables table `filter': Table does not exist (do you need to insmod?)
Perhaps ip6tables or your kernel needs to be upgraded.
+ true
+ ip6tables -A INPUT -j REJECT
ip6tables v1.6.0: can't initialize ip6tables table `filter': Table does not exist (do you need to insmod?)

```


https://www.qikqiak.com/post/encrypt-k8s-secrets-with-sealed-secrets/
