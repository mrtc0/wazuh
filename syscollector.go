package wazuh

import (
	"context"

	. "github.com/mrtc0/wazuh/define"
)

func GetHardwareInfoRequest(ctx context.Context, client httpClient, path string) (*GetHardwareInfoResponse, error) {
	response := &GetHardwareInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkAddrInfoRequest(ctx context.Context, client httpClient, path string) (*GetNetworkAddrInfoResponse, error) {
	response := &GetNetworkAddrInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkInterfaceInfoRequest(ctx context.Context, client httpClient, path string) (*GetNetworkInterfaceInfoResponse, error) {
	response := &GetNetworkInterfaceInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkProtocolInfoRequest(ctx context.Context, client httpClient, path string) (*GetNetworkProtocolInfoResponse, error) {
	response := &GetNetworkProtocolInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetOSInfoRequest(ctx context.Context, client httpClient, path string) (*GetOSInfoResponse, error) {
	response := &GetOSInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetPackagesInfoRequest(ctx context.Context, client httpClient, path string) (*GetPackagesInfoResponse, error) {
	response := &GetPackagesInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetPortsInfoRequest(ctx context.Context, client httpClient, path string) (*GetPortsInfoResponse, error) {
	response := &GetPortsInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetProcessesInfoRequest(ctx context.Context, client httpClient, path string) (*GetProcessesInfoResponse, error) {
	response := &GetProcessesInfoResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#id18
func (api *Client) GetHardwareInfo(agentID string) (*HardwareInfo, error) {
	return api.GetHardwareInfoContext(context.Background(), agentID)
}

func (api *Client) GetHardwareInfoContext(ctx context.Context, agentID string) (*HardwareInfo, error) {
	response, err := GetHardwareInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/hardware")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-address-info-of-an-agent
func (api *Client) GetNetworkAddrInfo(agentID string) (*[]NetworkAddrInfo, error) {
	return api.GetNetworkAddrInfoContext(context.Background(), agentID)
}

func (api *Client) GetNetworkAddrInfoContext(ctx context.Context, agentID string) (*[]NetworkAddrInfo, error) {
	response, err := GetNetworkAddrInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/netaddr")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-interface-info-of-an-agent
func (api *Client) GetNetworkInterfaceInfo(agentID string) (*[]NetworkInterfaceInfo, error) {
	return api.GetNetworkInterfaceInfoContext(context.Background(), agentID)
}

func (api *Client) GetNetworkInterfaceInfoContext(ctx context.Context, agentID string) (*[]NetworkInterfaceInfo, error) {
	response, err := GetNetworkInterfaceInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/netiface")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-protocol-info-of-an-agent
func (api *Client) GetNetworkProtocolInfo(agentID string) (*[]NetworkProtocolInfo, error) {
	return api.GetNetworkProtocolInfoContext(context.Background(), agentID)
}

func (api *Client) GetNetworkProtocolInfoContext(ctx context.Context, agentID string) (*[]NetworkProtocolInfo, error) {
	response, err := GetNetworkProtocolInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/netproto")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-os-info
func (api *Client) GetOSInfo(agentID string) (*OSInfo, error) {
	return api.GetOSInfoContext(context.Background(), agentID)
}

func (api *Client) GetOSInfoContext(ctx context.Context, agentID string) (*OSInfo, error) {
	response, err := GetOSInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/os")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-packages-info
func (api *Client) GetPackagesInfo(agentID string) (*[]PackagesInfo, error) {
	return api.GetPackagesInfoContext(context.Background(), agentID)
}

func (api *Client) GetPackagesInfoContext(ctx context.Context, agentID string) (*[]PackagesInfo, error) {
	response, err := GetPackagesInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/packages")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-ports-info-of-an-agent
func (api *Client) GetPortsInfo(agentID string) (*[]PortsInfo, error) {
	return api.GetPortsInfoContext(context.Background(), agentID)
}

func (api *Client) GetPortsInfoContext(ctx context.Context, agentID string) (*[]PortsInfo, error) {
	response, err := GetPortsInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/ports")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-processes-info
func (api *Client) GetProcessesInfo(agentID string) (*[]ProcessesInfo, error) {
	return api.GetProcessesInfoContext(context.Background(), agentID)
}

func (api *Client) GetProcessesInfoContext(ctx context.Context, agentID string) (*[]ProcessesInfo, error) {
	response, err := GetProcessesInfoRequest(ctx, api.httpclient, "syscollector/"+agentID+"/processes")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
