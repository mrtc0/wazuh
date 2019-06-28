package wazuh

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	httpclient httpClient
	Options    ClientOptions
}

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	Endpoint        url.URL
	certificatePath string
	keyPath         string
	certPem         []byte
	keyPem          []byte
	BasicUser       string
	BasicPass       string
}

func New(endpoint string, options ...ClientOption) (*Client, error) {
	var client *Client

	u, err := url.Parse(endpoint)
	if err != nil {
		return client, err
	}

	opt := ClientOptions{}
	for _, o := range options {
		o(&opt)
	}

	opt.Endpoint = *u

	tr, err := setCertficate(opt)

	if err != nil {
		return client, err
	}

	client = &Client{
		httpclient: &http.Client{Transport: tr},
		Options:    opt,
	}

	return client, nil
}

func WithBasicAuth(username, password string) ClientOption {
	return func(options *ClientOptions) {
		options.BasicUser = username
		options.BasicPass = password
	}
}

// If you are doing client certificate authentication, use it
func WithClientCertificateFromFile(cert string, key string) ClientOption {
	return func(options *ClientOptions) {
		options.certificatePath = cert
		options.keyPath = key
	}
}

func WithClientCertificate(certPEMBlock, keyPEMBlock []byte) ClientOption {
	return func(options *ClientOptions) {
		options.certPem = certPEMBlock
		options.keyPem = keyPEMBlock
	}
}

func setCertficate(opt ClientOptions) (*http.Transport, error) {
	var cert tls.Certificate
	var tr *http.Transport
	var err error

	if opt.certificatePath != "" && opt.keyPath != "" {
		cert, err = tls.LoadX509KeyPair(opt.certificatePath, opt.keyPath)

		if err != nil {
			return tr, err
		}
	}

	if len(opt.certPem) > 0 && len(opt.keyPem) > 0 {
		cert, err = tls.X509KeyPair(opt.certPem, opt.keyPem)
		if err != nil {
			return tr, err
		}
	}

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}},
	}

	return tr, nil
}
