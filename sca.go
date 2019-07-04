package wazuh

import (
	"context"
	"path"

	. "github.com/mrtc0/wazuh/define/sca"
)

// Security Configuration Assessment
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#security-configuration-assessment

// Get security configuration assessment (SCA) checks database
// Returns the sca checks of an agent
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#id19
func (client *Client) GetScaChecks(agentId, Id string) (*[]ScaCheck, error) {
	return client.GetScaChecksContext(context.Background(), agentId, Id)
}

func (client *Client) GetScaChecksContext(ctx context.Context, agentId, Id string) (*[]ScaCheck, error) {
	response, err := GetScaChecksRequest(ctx, client, path.Join("sca", agentId, "checks", Id))

	if err != nil {
		return nil, err
	}

	return &response.Data.Items, nil
}

func GetScaChecksRequest(ctx context.Context, client *Client, path string) (*GetScaChecksResponse, error) {
	response := &GetScaChecksResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get security configuration assessment (SCA) database
// Returns the sca database of an agent.
// https://documentation.wazuh.com/3.x/user-manual/api/reference.html#get-security-configuration-assessment-sca-database
func (client *Client) GetSca(agentId string) (*[]Sca, error) {
	return client.GetScaContext(context.Background(), agentId)
}

func (client *Client) GetScaContext(ctx context.Context, agentId string) (*[]Sca, error) {
	response, err := GetScaRequest(ctx, client, path.Join("sca", agentId))

	if err != nil {
		return nil, err
	}

	return &response.Data.Items, nil
}

func GetScaRequest(ctx context.Context, client *Client, path string) (*GetScaResponse, error) {
	response := &GetScaResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
