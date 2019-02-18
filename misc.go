package wazuh

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

func GetJson(ctx context.Context, client httpClient, endpoint string, intf interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	if USERNAME != "" || PASSWORD != "" {
		req.SetBasicAuth(USERNAME, PASSWORD)
	}

	/*
		q := req.URL.Query()
		q.Add("api_key", "key_from_environment_or_flag")
		req.URL.RawQuery = q.Encode()
	*/

	return DoGet(ctx, client, req, intf)
}
