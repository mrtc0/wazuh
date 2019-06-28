package wazuh

import (
	"context"
	"path"

	. "github.com/mrtc0/wazuh/define/decoder"
)

func GetAllDecodersRequest(ctx context.Context, client *Client, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-decoders
func (client *Client) GetAllDecoders() (*[]Decoder, error) {
	return client.GetAllDecodersContext(context.Background())
}

func (client *Client) GetAllDecodersContext(ctx context.Context) (*[]Decoder, error) {
	response, err := GetAllDecodersRequest(ctx, client, "decoders")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetDecoderByNameRequest(ctx context.Context, client *Client, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-decoders-by-name
func (client *Client) GetDecoderByName(name string) (*[]Decoder, error) {
	return client.GetDecoderByNameContext(context.Background(), name)
}

func (client *Client) GetDecoderByNameContext(ctx context.Context, name string) (*[]Decoder, error) {
	response, err := GetDecoderByNameRequest(ctx, client, path.Join("decoders", name))

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetDecoderFilesRequest(ctx context.Context, client *Client, path string) (*GetDecoderFilesResponse, error) {
	response := &GetDecoderFilesResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-decoders-files
func (client *Client) GetDecoderFiles() (*[]DecoderFiles, error) {
	return client.GetDecoderFilesContext(context.Background())
}

func (client *Client) GetDecoderFilesContext(ctx context.Context) (*[]DecoderFiles, error) {
	response, err := GetDecoderFilesRequest(ctx, client, path.Join("decoders", "files"))

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetParentDecodersRequest(ctx context.Context, client *Client, path string) (*GetDecodersResponse, error) {
	response := &GetDecodersResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-parent-decoders
func (client *Client) GetParentDecoders() (*[]Decoder, error) {
	return client.GetParentDecodersContext(context.Background())
}

func (client *Client) GetParentDecodersContext(ctx context.Context) (*[]Decoder, error) {
	response, err := GetParentDecodersRequest(ctx, client, path.Join("decoders", "parents"))

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
