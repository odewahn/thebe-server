package main

import (
    "github.com/samalba/dockerclient"
    "log"
    "fmt"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
)

// Reads the specified file into a byte array
func fetchFile(path, fn string) ([]byte, error) {
  fileName := fmt.Sprintf("%s/%s", path, fn)
  out, err := ioutil.ReadFile(fileName)
  if err != nil {
    log.Fatalf("error loading %s : %s", fn, err)
  }
  return out, err
}

// Loads and returns a tls config settings for the swarm.  Lightly adapted version of:
//     https://github.com/ehazlett/interlock/blob/master/interlock/main.go#L14-L32
func getTLSConfig(certsDir string) (*tls.Config, error) {
	// TLS config
	var tlsConfig tls.Config

  caCert, err := fetchFile(certsDir, "ca.pem")
  cert, err := fetchFile(certsDir, "cert.pem")
  key, err := fetchFile(certsDir, "key.pem")

	tlsConfig.InsecureSkipVerify = true
	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = certPool
	keypair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{keypair}

	return &tlsConfig, nil
}


func main() {
    // Init the client
    dockerURL := "tcp://104.130.0.52:2376"

    certsDir := "/Users/apple/Desktop/9fadfa89-0400-453a-a7eb-436aea737831"

    tlsConfig, err := getTLSConfig(certsDir)

    docker, _ := dockerclient.NewDockerClient(dockerURL, tlsConfig)

    // Get only running containers
    containers, err := docker.ListContainers(false, false, "")
    if err != nil {
        log.Fatal(err)
    }
    for _, c := range containers {
        log.Println(c.Id, c.Names)
    }


}
