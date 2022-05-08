package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"go.etcd.io/etcd/clientv3"
	"log"
	"os"
)

const (
	DEFAULT_ETCD_ENDPOINT = "127.0.0.1:2379"
	PATH_PREFIX           = "etcd://"
)

func main() {

	if len(os.Args) != 5 {
		log.Fatalf("invalid input parameter: %v", os.Args)
	}

	var cli *clientv3.Client
	var tlsConfig *tls.Config
	var r *clientv3.GetResponse
	var err error

	if tlsConfig, err = loadTLS(); err != nil {
		log.Fatalf("failed to read tls configuration: %v", os.Args)
	}

	if cli, err = clientv3.New(
		clientv3.Config{
			Endpoints: []string{DEFAULT_ETCD_ENDPOINT},
			TLS:       tlsConfig,
		},
	); err != nil {
		log.Fatalf("failed to create etcd client: %v", err)
	}

	defer cli.Close()

	if !strings.HasPrefix(os.Args[4], PATH_PREFIX) {
		log.Fatalf("malformed path: %v", os.Args[4])
	}
	key := strings.Replace(os.Args[4], PATH_PREFIX, "", -1)
	if r, err = cli.Get(context.Background(), key); err != nil {
		log.Fatalf("failed to read value: %v", err)
	}
	for _, kv := range r.Kvs {
		v, err := base64.StdEncoding.DecodeString(string(kv.Value))
		if err != nil {
			log.Fatalf("values are not base64 encoded: %s", string(kv.Value))
		}
		fmt.Print(string(v))
	}

}

func loadTLS() (*tls.Config, error) {

	configFiles := path.Join(os.Getenv("HELM_PLUGIN_DIR"), "config.json")

	r, err := ioutil.ReadFile(configFiles)
	if err != nil {
		return nil, err
	}
	var config map[string]string
	if err := json.Unmarshal(r, &config); err != nil {
		return nil, err
	}

	caFile := config["caFile"]
	certFile := config["certFile"]
	keyFile := config["keyFile"]

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	caData, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}

	return &tlsConfig, nil
}
