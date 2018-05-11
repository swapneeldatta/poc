package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const serverCertificate = "/home/buga/eclipse-workspace/MutualAuth/developer.cer"
const clientCertificate = "/home/buga/eclipse-workspace/MutualAuth/developer.cer"
const clientKey = "/home/buga/eclipse-workspace/MutualAuth/developer.p12"

func main() {
	client := InitializeHTTPClient()
	log.Println("Hello! Initializing...")
	go func() {
		for {
			r, e := client.Get("https://localhost:8443/app/poc/ping")
			log.Println(r, e)
			time.Sleep(1 * time.Second)
		}
	}()
}

// InitializeHTTPClient initializes an HTTP Connection
func InitializeHTTPClient() *http.Client {
	return InitializeHTTPClientWithTimeout(1)
}

// InitializeHTTPClientWithTimeout initializes an HTTP Connection with timeout
func InitializeHTTPClientWithTimeout(timeout uint) *http.Client {
	certificates, errors := ioutil.ReadFile(serverCertificate)
	HandleError("attempting to load server certificates", &errors, true)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certificates)

	cert, errors := tls.LoadX509KeyPair(clientCertificate, clientKey)
	HandleError("attempting to load client certificates", &errors, true)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		Timeout: time.Duration(timeout) * time.Second,
	}
	return client
}

// PostWithAPIKey does as the name suggests
func PostWithAPIKey(c *http.Client, apiKey *string, url *string, data *string) (*http.Response, error) {
	request, e := http.NewRequest("POST", *url, bytes.NewBufferString(*data))
	HandleError("creating post request with apikey", &e, false)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "DATTUS_AUTH")
	request.Header.Add("api-key", *apiKey)
	response, e := c.Do(request)
	return response, e
}

// HandleError encapsulates a try catch block with a nice log output
func HandleError(attemptingToDoWhat string, e *error, failOnError bool) {
	if *e != nil {
		log.Println("Error", attemptingToDoWhat, "(", *e, ")")
		if failOnError {
			log.Panic("Terminating application because of previous errors")
		}
	}
}
