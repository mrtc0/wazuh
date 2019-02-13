package wazuh

import (
	"crypto/tls"
	"net/http"
)

var APIURL string
var USERNAME string
var PASSWORD string
var CLIENT_CERT_PATH string
var CLIENT_KEY_PATH string

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	httpclient httpClient
}

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	certificatePath string
	keyPath string
}

func New(url string, options ...ClientOption) *Client {
	APIURL = url

	opt := ClientOptions{}
	for _, o := range options {
		o(&opt)
	}

	cert, _ := tls.LoadX509KeyPair(opt.certificatePath, opt.keyPath)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}},
	}

	s := &Client{
		httpclient: &http.Client{Transport: tr},
	}

	return s
}

func (api *Client) SetBasicAuth(username, password string) {
	USERNAME = username
	PASSWORD = password
}

func SetClientCertificate(cert string, key string) ClientOption {
	CLIENT_CERT_PATH = cert
	CLIENT_KEY_PATH = key

	return func(options *ClientOptions) {
		options.certificatePath = cert
		options.keyPath = key
	}
}
