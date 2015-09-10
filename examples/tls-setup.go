package main

import (
    "github.com/samalba/dockerclient"
    "log"
    "fmt"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
)

func getTLSConfig(caCert, cert, key []byte, allowInsecure bool) (*tls.Config, error) {
	// TLS config
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = true
	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = certPool
	keypair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{keypair}
	if allowInsecure {
		tlsConfig.InsecureSkipVerify = true
	}

	return &tlsConfig, nil
}


func main() {
    // Init the client
    docker_url := "tcp://104.130.0.52:2376"

    certs_dir := "/Users/apple/Desktop/9fadfa89-0400-453a-a7eb-436aea737831"

    caCertFile := fmt.Sprintf("%s/ca.pem", certs_dir)
    caCert, err := ioutil.ReadFile(caCertFile)
    if err != nil {
      log.Fatalf("error loading tls ca cert: %s", err)
    }

    certFile := fmt.Sprintf("%s/cert.pem", certs_dir)
    cert, err := ioutil.ReadFile(certFile)
    if err != nil {
      log.Fatalf("error loading tls cert: %s", err)
    }

    keyFile := fmt.Sprintf("%s/key.pem", certs_dir)
    key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			log.Fatalf("error loading tls key: %s", err)
		}


    tlsConfig, err := getTLSConfig(caCert, cert, key, false)

    docker, _ := dockerclient.NewDockerClient(docker_url, tlsConfig)

    // Get only running containers
    containers, err := docker.ListContainers(false, false, "")
    if err != nil {
        log.Fatal(err)
    }
    for _, c := range containers {
        log.Println(c.Id, c.Names)
    }


}
