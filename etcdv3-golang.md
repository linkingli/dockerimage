```
import (
    "crypto/tls"
    "crypto/x509"
    etcdclient "github.com/coreos/etcd/clientv3"
)
func NewEtcdCluster(etcdAdds []string, keyStatusPrifix, keyConfigPrifix, configDir string) (*EtcdCluster, error) {
 
var etcdCertPath = configDir + "/etcd/client.pem"
    var etcdCertKeyPath = configDir + "/etcd/client-key.pem"
    var etcdCaPath = configDir + "/etcd/ca.pem"
 
    // load cert
    cert, err := tls.LoadX509KeyPair(etcdCertPath, etcdCertKeyPath)
    if err != nil {
        return nil, err
    }
 
    // load root ca
    caData, err := ioutil.ReadFile(etcdCaPath)
    if err != nil {
        return nil, err
    }
 
    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM(caData)
 
    _tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs:      pool,
    }
 
    cfg := etcdclient.Config{
        Endpoints: etcdAdds,
        TLS:       _tlsConfig,
    }
    client, err := etcdclient.New(cfg)
    if err != nil {
        return nil, err
    }
 
    cluster := &EtcdCluster{
        keyStatusPrifix: keyStatusPrifix,
        keyConfigPrifix: keyConfigPrifix,
        etcdAdds:        etcdAdds,
        ticker:          time.NewTicker(time.Second * 5),
        nodes:           make(map[uint16]INodeStatus),
        kapi:            client,
    }
    return cluster, err
}
```
