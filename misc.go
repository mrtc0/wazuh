package wazuh

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_path "path"
)

func ParseResponseBody(body io.ReadCloser, intf interface{}) error {
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(response, intf)
}

func DoGet(ctx context.Context, client httpClient, req *http.Request, intf interface{}) error {
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 < resp.StatusCode {
		return fmt.Errorf("http request error code:%d msg: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	return ParseResponseBody(resp.Body, intf)
}

func GetJson(ctx context.Context, client *Client, path string, intf interface{}) error {
	endpoint := client.Options.Endpoint
	endpoint.Path = _path.Join(endpoint.Path, path)

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return err
	}

	if client.Options.BasicUser != "" || client.Options.BasicPass != "" {
		req.SetBasicAuth(client.Options.BasicUser, client.Options.BasicPass)
	}

	/*
		q := req.URL.Query()
		q.Add("api_key", "key_from_environment_or_flag")
		req.URL.RawQuery = q.Encode()
	*/

	return DoGet(ctx, client.httpclient, req, intf)
}
