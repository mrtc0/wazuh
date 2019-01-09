package wazuh

import (
	"context"

	. "github.com/mrtc0/wazuh/define/decoder"
)

func GetAllDecodersRequest(ctx context.Context, client httpClient, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-decoders
func (api *Client) GetAllDecoders() (*[]Decoder, error) {
	return api.GetAllDecodersContext(context.Background())
}

func (api *Client) GetAllDecodersContext(ctx context.Context) (*[]Decoder, error) {
	response, err := GetAllDecodersRequest(ctx, api.httpclient, "decoders")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetDecoderByNameRequest(ctx context.Context, client httpClient, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-decoders-by-name
func (api *Client) GetDecoderByName(name string) (*[]Decoder, error) {
	return api.GetDecoderByNameContext(context.Background(), name)
}

func (api *Client) GetDecoderByNameContext(ctx context.Context, name string) (*[]Decoder, error) {
	response, err := GetDecoderByNameRequest(ctx, api.httpclient, "decoders/"+name+"")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetDecoderFilesRequest(ctx context.Context, client httpClient, path string) (*GetDecoderFilesResponse, error) {
	response := &GetDecoderFilesResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-decoders-files
func (api *Client) GetDecoderFiles() (*[]DecoderFiles, error) {
	return api.GetDecoderFilesContext(context.Background())
}

func (api *Client) GetDecoderFilesContext(ctx context.Context) (*[]DecoderFiles, error) {
	response, err := GetDecoderFilesRequest(ctx, api.httpclient, "decoders/files")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetParentDecodersRequest(ctx context.Context, client httpClient, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-parent-decoders
func (api *Client) GetParentDecoders() (*[]Decoder, error) {
	return api.GetParentDecodersContext(context.Background())
}

func (api *Client) GetParentDecodersContext(ctx context.Context) (*[]Decoder, error) {
	response, err := GetParentDecodersRequest(ctx, api.httpclient, "decoders/parents")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
