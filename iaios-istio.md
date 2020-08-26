```
yanganbao KF8394 (Partner)	第一个就是上面那个创建[CREATE_VIRTUAL]报错，需要把istio-galley这个删掉，并且把ValidatingWebhookConfiguration删除才能创建成功 
	14:13
	第二个问题是创建RC的时候报  Error creating: Internal error occurred: failed calling admission webhook "sidecar-injector.istio.io": Post https://istio-sidecar-injector.cloudos-mesh.svc:443/inject?timeout=30s: x509: certificate signed by unknown authority实
	

```
