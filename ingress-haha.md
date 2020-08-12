```
apiVersion:  extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-model-serving
  namespace: cloudos-hpai-system
  annotations:
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
    - http:
        paths:
          - path: /kernel
            backend:
              serviceName: kernel-service
              servicePort: 8088

```
