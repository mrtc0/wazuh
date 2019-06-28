package wazuh_test

import (
	"log"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/mrtc0/wazuh"
)

var (
	serverAddr string
	endpoint   string
	once       sync.Once
)

func startServer() {
	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	endpoint = "http://" + serverAddr + "/"
	log.Print("Test server listening on ", serverAddr)
}

func TestCreateClientWithBasicAuth(t *testing.T) {
	wazuh, err := wazuh.New("https://wazuh.localhost:55000/", wazuh.WithBasicAuth("username", "password"))
	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if wazuh.Options.BasicUser != "username" {
		t.Fatalf("expect username, but actual %#v", wazuh.Options.BasicUser)
	}

	if wazuh.Options.BasicPass != "password" {
		t.Fatalf("expect password, but actual %#v", wazuh.Options.BasicPass)
	}
}

func ExampleNew() {
	wazuh.New("https://wazuh.localhost:55000/")
}

func ExampleNewWithBasicAuth() {
	wazuh.New("https://wazuh.localhost:55000/", wazuh.WithBasicAuth("username", "password"))
}

func ExampleNewWithClientCertificateFromFile() {
	wazuh.New("https://wazuh.localhost:55000/", wazuh.WithClientCertificateFromFile("/path/to/certificate.cert", "/path/to/private.key"))
}

func ExampleNewWithClientCertificate() {
	var cert, key []byte
	wazuh.New("https://wazuh.localhost:55000/", wazuh.WithClientCertificate(cert, key))
}
