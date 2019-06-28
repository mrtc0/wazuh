package wazuh

import (
	"context"
	"path"

	. "github.com/mrtc0/wazuh/define"
)

func GetHardwareInfoRequest(ctx context.Context, client *Client, path string) (*GetHardwareInfoResponse, error) {
	response := &GetHardwareInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkAddrInfoRequest(ctx context.Context, client *Client, path string) (*GetNetworkAddrInfoResponse, error) {
	response := &GetNetworkAddrInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkInterfaceInfoRequest(ctx context.Context, client *Client, path string) (*GetNetworkInterfaceInfoResponse, error) {
	response := &GetNetworkInterfaceInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetNetworkProtocolInfoRequest(ctx context.Context, client *Client, path string) (*GetNetworkProtocolInfoResponse, error) {
	response := &GetNetworkProtocolInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetOSInfoRequest(ctx context.Context, client *Client, path string) (*GetOSInfoResponse, error) {
	response := &GetOSInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetPackagesInfoRequest(ctx context.Context, client *Client, path string) (*GetPackagesInfoResponse, error) {
	response := &GetPackagesInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetPortsInfoRequest(ctx context.Context, client *Client, path string) (*GetPortsInfoResponse, error) {
	response := &GetPortsInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetProcessesInfoRequest(ctx context.Context, client *Client, path string) (*GetProcessesInfoResponse, error) {
	response := &GetProcessesInfoResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#id18
func (client *Client) GetHardwareInfo(agentID string) (*HardwareInfo, error) {
	return client.GetHardwareInfoContext(context.Background(), agentID)
}

func (client *Client) GetHardwareInfoContext(ctx context.Context, agentID string) (*HardwareInfo, error) {
	response, err := GetHardwareInfoRequest(ctx, client, path.Join("syscollector", agentID, "hardware"))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-address-info-of-an-agent
func (client *Client) GetNetworkAddrInfo(agentID string) (*[]NetworkAddrInfo, error) {
	return client.GetNetworkAddrInfoContext(context.Background(), agentID)
}

func (client *Client) GetNetworkAddrInfoContext(ctx context.Context, agentID string) (*[]NetworkAddrInfo, error) {
	response, err := GetNetworkAddrInfoRequest(ctx, client, path.Join("syscollector", agentID, "netaddr"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-interface-info-of-an-agent
func (client *Client) GetNetworkInterfaceInfo(agentID string) (*[]NetworkInterfaceInfo, error) {
	return client.GetNetworkInterfaceInfoContext(context.Background(), agentID)
}

func (client *Client) GetNetworkInterfaceInfoContext(ctx context.Context, agentID string) (*[]NetworkInterfaceInfo, error) {
	response, err := GetNetworkInterfaceInfoRequest(ctx, client, path.Join("syscollector/", agentID, "netiface"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-network-protocol-info-of-an-agent
func (client *Client) GetNetworkProtocolInfo(agentID string) (*[]NetworkProtocolInfo, error) {
	return client.GetNetworkProtocolInfoContext(context.Background(), agentID)
}

func (client *Client) GetNetworkProtocolInfoContext(ctx context.Context, agentID string) (*[]NetworkProtocolInfo, error) {
	response, err := GetNetworkProtocolInfoRequest(ctx, client, path.Join("syscollector", agentID, "netproto"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-os-info
func (client *Client) GetOSInfo(agentID string) (*OSInfo, error) {
	return client.GetOSInfoContext(context.Background(), agentID)
}

func (client *Client) GetOSInfoContext(ctx context.Context, agentID string) (*OSInfo, error) {
	response, err := GetOSInfoRequest(ctx, client, path.Join("syscollector", agentID, "os"))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-packages-info
func (client *Client) GetPackagesInfo(agentID string) (*[]PackagesInfo, error) {
	return client.GetPackagesInfoContext(context.Background(), agentID)
}

func (client *Client) GetPackagesInfoContext(ctx context.Context, agentID string) (*[]PackagesInfo, error) {
	response, err := GetPackagesInfoRequest(ctx, client, path.Join("syscollector", agentID, "packages"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-ports-info-of-an-agent
func (client *Client) GetPortsInfo(agentID string) (*[]PortsInfo, error) {
	return client.GetPortsInfoContext(context.Background(), agentID)
}

func (client *Client) GetPortsInfoContext(ctx context.Context, agentID string) (*[]PortsInfo, error) {
	response, err := GetPortsInfoRequest(ctx, client, path.Join("syscollector", agentID, "ports"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-processes-info
func (client *Client) GetProcessesInfo(agentID string) (*[]ProcessesInfo, error) {
	return client.GetProcessesInfoContext(context.Background(), agentID)
}

func (client *Client) GetProcessesInfoContext(ctx context.Context, agentID string) (*[]ProcessesInfo, error) {
	response, err := GetProcessesInfoRequest(ctx, client, path.Join("syscollector", agentID, "processes"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
