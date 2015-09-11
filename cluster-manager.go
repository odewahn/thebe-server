package main

import (
    "github.com/samalba/dockerclient"
    "log"
    "fmt"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "os"
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



func launchNotebook(docker *dockerclient.DockerClient, hostname, image string) (error) {

    // Create a container
    containerConfig := &dockerclient.ContainerConfig{
        Image: image,
        Cmd:   []string{"/bin/sh", "-c", "ipython notebook --ip=0.0.0.0 --no-browser"},
        ExposedPorts: map[string]struct{}{
          "8888/tcp": {},
        },
        Hostname: hostname,
        Domainname: os.Getenv("THEBE_SERVER_BASE_URL"),
    }

    containerId, err := docker.CreateContainer(containerConfig, hostname)
    if err != nil {
        log.Println(err)
    }

    // Start the container
    hostConfig := &dockerclient.HostConfig{
      PublishAllPorts: true,
    }
    err = docker.StartContainer(containerId, hostConfig)
    if err != nil {
        log.Println(err)
    }
    return err

}
