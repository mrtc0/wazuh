package wazuh

import (
	"crypto/tls"
	"net/http"
)

var APIURL string
var USERNAME string
var PASSWORD string

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	httpclient httpClient
}

func New(url string) *Client {
	APIURL = url
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
