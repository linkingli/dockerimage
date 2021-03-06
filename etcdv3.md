```
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	// 为了保证 HTTPS 链接可信，需要预先加载目标证书签发机构的 CA 根证书
	etcdCA, err := ioutil.ReadFile("/Users/mritd/tmp/etcd_ssl/etcd-root-ca.pem")
	if err != nil {
		log.Fatal(err)
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair("/Users/mritd/tmp/etcd_ssl/etcd.pem", "/Users/mritd/tmp/etcd_ssl/etcd-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个空的 CA Pool
	// 因为后续只会链接 Etcd 的 api 端点，所以此处选择使用空的 CA Pool，然后只加入 Etcd CA 既可
	// 如果期望链接其他 TLS 端点，那么最好使用 x509.SystemCertPool() 方法先 copy 一份系统根 CA
	// 然后再向这个 Pool 中添加自定义 CA
	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(etcdCA)

	// 创建 api v3 的 client
	cli, err := clientv3.New(clientv3.Config{
		// etcd https api 端点
		Endpoints:   []string{"https://172.16.14.114:2379"},
		DialTimeout: 5 * time.Second,
		// 自定义 CA 及 Client Cert 配置
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	putResp, err := cli.Put(ctx, "sample_key", "sample_value")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(putResp)
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	delResp, err := cli.Delete(ctx, "sample_key")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(delResp)
	}
	cancel()
}
```
```
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "time"

    "crypto/tls"
    "crypto/x509"

    "go.etcd.io/etcd/clientv3"
    "golang.org/x/net/context"
)

var (
    dialTimeout    = 5 * time.Second
    requestTimeout = 4 * time.Second
    endpoints      = []string{"https://172.17.84.204:2379", "https://172.17.84.205:2379", "https://172.17.84.206:2379"}
)

func main() {

    var etcdCert = "./ca/etcd-client.pem"
    var etcdCertKey = "./ca/etcd-client-key.pem"
    var etcdCa = "./ca/ca.pem"

    cert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
    if err != nil {
        return
    }

    caData, err := ioutil.ReadFile(etcdCa)
    if err != nil {
        return
    }

    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM(caData)

    _tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs:      pool,
    }

    cfg := clientv3.Config{
        Endpoints: endpoints,
        TLS:       _tlsConfig,
    }

    cli, err := clientv3.New(cfg)

    if err != nil {
        log.Fatal(err)
    }

    defer cli.Close()

    key1, value1 := "testkey1", "value"

    ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
    _, err = cli.Put(ctx, key1, value1)
    cancel()
    if err != nil {
        log.Println("Put failed. ", err)
    } else {
        log.Printf("Put {%s:%s} succeed\n", key1, value1)
    }

    ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
    resp, err := cli.Get(ctx, key1)
    cancel()
    if err != nil {
        log.Println("Get failed. ", err)
        return
    }

    for _, kv := range resp.Kvs {
        log.Printf("Get {%s:%s} \n", kv.Key, kv.Value)
    }

    done := make(chan bool)

    go func() {
        wch := cli.Watch(context.Background(), key1)

        for item := range wch {
            for _, ev := range item.Events {
                log.Printf("Type:%s, key:%s, value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
            }
        }
    }()

    go func() {
        for cnt := 0; cnt < 10; cnt++ {
            value := fmt.Sprintf("%s%d", "value", cnt)
            _, err = cli.Put(context.Background(), key1, value)
            if err != nil {
                log.Println("Put failed. ", err)
            } else {
                log.Printf("Put {%s:%s} succeed\n", key1, value)
            }
        }
    }()

    <-done

    log.Println("Done!")
}
```
