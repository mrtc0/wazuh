package wazuh

import (
	"log"
	"net/http/httptest"
	"sync"
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
