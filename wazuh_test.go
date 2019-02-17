package wazuh_test

import (
	"log"
	"net/http/httptest"
	"sync"

	"github.com/mrtc0/wazuh"
)

var (
	serverAddr string
	once       sync.Once
)

func startServer() {
	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	log.Print("Test server listening on ", serverAddr)
}

func ExampleNew() {
	wazuh.New("https://wazuh.localhost:55000/")
}

func ExampleNew_withBasicAuth() {
	wazuh := wazuh.New("https://wazuh.localhost:55000/")
	// When using Basic Auth
	wazuh.SetBasicAuth("username", "password")
}

func ExampleNew_withClientCertificate() {
	// When using Client Certificate
	wazuh.New("https://wazuh.localhost:55000/", wazuh.SetClientCertificate("/path/to/certificate.cert", "/path/to/private.key"))
}
