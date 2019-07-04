package wazuh

import (
	"context"
	"path"

	. "github.com/mrtc0/wazuh/define/syscheck"
)

// Get syscheck files
// Returns the syscheck files of an agent.
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#get-syscheck-files
func (client *Client) GetSyscheckFiles(agentId string) (*[]SyscheckFiles, error) {
	return client.GetSyscheckFilesContext(context.Background(), agentId)
}

func (client *Client) GetSyscheckFilesContext(ctx context.Context, agentId string) (*[]SyscheckFiles, error) {
	response, err := GetSyscheckFilesRequest(ctx, client, path.Join("syscheck", agentId))

	if err != nil {
		return nil, err
	}

	return &response.Data.Items, nil
}

func GetSyscheckFilesRequest(ctx context.Context, client *Client, path string) (*GetSyscheckFilesResponse, error) {
	response := &GetSyscheckFilesResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Run sys check scan in all agents
// Runs syscheck and rootcheck on all agents (Wazuh launches both processes simultaneously).
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#id22
func (client *Client) RunSyscheckAllAgents() (*string, error) {
	return client.RunSyscheckAllAgentsContext(context.Background())
}

func (client *Client) RunSyscheckAllAgentsContext(ctx context.Context) (*string, error) {
	response, err := RunSyscheckAllAgentsRequest(ctx, client, path.Join("syscheck"))

	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func RunSyscheckAllAgentsRequest(ctx context.Context, client *Client, path string) (*RunSyscheckAllAgentsResponse, error) {
	response := &RunSyscheckAllAgentsResponse{}
	err := PutJson(ctx, client, path, []byte{}, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Run syscheck scan in an agent
// Runs syscheck and rootcheck on an agent (Wazuh launches both processes simultaneously).
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#run-syscheck-scan-in-an-agent
func (client *Client) RunSyscheckAgent(agentId string) (*string, error) {
	return client.RunSyscheckAgentContext(context.Background(), agentId)
}

func (client *Client) RunSyscheckAgentContext(ctx context.Context, agentId string) (*string, error) {
	response, err := RunSyscheckAgentRequest(ctx, client, path.Join("syscheck", agentId))

	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func RunSyscheckAgentRequest(ctx context.Context, client *Client, path string) (*RunSyscheckAgentResponse, error) {
	response := &RunSyscheckAgentResponse{}
	err := PutJson(ctx, client, path, []byte{}, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get last syscheck scan
// Return the timestamp of the last syscheck scan.
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#get-last-syscheck-scan
func (client *Client) GetLastSyscheckScan(agentId string) (*LastSyscheckScan, error) {
	return client.GetLastSyscheckScanContext(context.Background(), agentId)
}

func (client *Client) GetLastSyscheckScanContext(ctx context.Context, agentId string) (*LastSyscheckScan, error) {
	response, err := GetLastSyscheckScanRequest(ctx, client, path.Join("syscheck", agentId, "last_scan"))

	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func GetLastSyscheckScanRequest(ctx context.Context, client *Client, path string) (*GetLastSyscheckScanResponse, error) {
	response := &GetLastSyscheckScanResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Clear syscheck database of an agent
// Clears the syscheck database for the specified agent.
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#id20
func (client *Client) ClearSyscheckDatabase(agentId string) (*string, error) {
	return client.ClearSyscheckDatabaseContext(context.Background(), agentId)
}

func (client *Client) ClearSyscheckDatabaseContext(ctx context.Context, agentId string) (*string, error) {
	response, err := ClearSyscheckDatabaseRequest(ctx, client, path.Join("syscheck", agentId))

	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func ClearSyscheckDatabaseRequest(ctx context.Context, client *Client, path string) (*ClearSyscheckDatabaseResponse, error) {
	response := &ClearSyscheckDatabaseResponse{}
	err := DoDelete(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
