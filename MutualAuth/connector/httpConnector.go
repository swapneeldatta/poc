package connector

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"dattus.com/errorhandler"
)

const serverCertificate = "/etc/ssl/certs/dattusPortal.crt"
const clientCertificate = "/etc/ssl/private/dattusClient.crt"
const clientKey = "/etc/ssl/private/dattusClient.key"

// InitializeHTTPClient initializes an HTTP Connection
func InitializeHTTPClient() *http.Client {
	return InitializeHTTPClientWithTimeout(1)
}

// InitializeHTTPClientWithTimeout initializes an HTTP Connection with timeout
func InitializeHTTPClientWithTimeout(timeout uint) *http.Client {
	certificates, errors := ioutil.ReadFile(serverCertificate)
	errorhandler.HandleError("attempting to load server certificates", &errors, true)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certificates)

	cert, errors := tls.LoadX509KeyPair(clientCertificate, clientKey)
	errorhandler.HandleError("attempting to load client certificates", &errors, true)

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
	errorhandler.HandleError("creating post request with apikey", &e, false)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "DATTUS_AUTH")
	request.Header.Add("api-key", *apiKey)
	response, e := c.Do(request)
	return response, e
}
