```
go build -o alert-config -gcflags=-trimpath=/go/src -ldflags '-s -w -X github.com/yylt/prometheus-operator/api.Version=b830db4 
-X github.com/yylt/prometheus-operator/api.Revision=go1.12.9 -X github.com/yylt/prometheus-operator/api.Dirty=dirty
-X github.com/yylt/prometheus-operator/api.Branch=(HEAD' github.com/yylt/prometheus-operator
```
