```
require(k8s.io/client-go v0.16.8)
```
```
import (
	"fmt"
	"io/ioutil"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/cert"
	"log"
)

func getClient(tokenPath, caPath, masterUrl string) *kubernetes.Clientset {
	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		fmt.Println(err)
	}

	tlsClientConfig := rest.TLSClientConfig{}

	if _, err := cert.NewPool(caPath); err != nil {
		log.Printf("Expected to load root CA config from %s, but got err: %v", caPath, err)
	} else {
		tlsClientConfig.CAFile = caPath
	}
	config := rest.Config{
		Host:            masterUrl,
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
		BearerTokenFile: tokenPath,
	}
	client, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Printf("faild to create kuber client %s", err)
	}
	return client
}

func main() {
	clientset := getClient("D:\\token", "D:\\ca.crt", "10.125.31.171:8443")
	pods, err := clientset.CoreV1().Pods("default").List(v1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}
}
```
