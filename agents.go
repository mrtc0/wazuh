package wazuh

import (
	"context"
	"path"
)

type AgentInformationData struct {
	TotalItems int                `json:"totalItems"`
	Items      []AgentInformation `json:"items"`
}

type GetAllAgentsResponse struct {
	Error int                  `json:"error"`
	Data  AgentInformationData `json:"data"`
}

type GetAnAgentResponse struct {
	Error int              `json:"error"`
	Data  AgentInformation `json::"data"`
}

type GetGroupsResponse struct {
	Error int       `json:"error"`
	Data  GroupData `json:"data"`
}

type AgentRestartResponse struct {
	Error int              `json:"error"`
	Data  AgentRestartData `json:"data"`
}

type AgentRestartData struct {
	Msg            string   `json:"msg"`
	AffectedAgents []string `json:"affectedAgents"`
}

type GroupData struct {
	TotalItems int     `json:"totalItems"`
	Items      []Group `json:"items"`
}

type Group struct {
	Count     int    `json:"count"`
	MergedSum string `json:"mergedSum"`
	ConfigSum string `json:"configSum"`
	Name      string `json:"name"`
}

type Os struct {
	Major    string `json:"major"`
	Name     string `json:"name"`
	Uname    string `json:"uname"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Codename string `json:"codename"`
	Arch     string `json:"arch"`
	Minor    string `json:"minor"`
}

type AgentInformation struct {
	Status        string   `json:"status"`
	Name          string   `json:"name"`
	IP            string   `json:"ip"`
	Manager       string   `json:"manager"`
	NodeName      string   `json:"node_name"`
	DateAdd       string   `json:"dateAdd"`
	Version       string   `json:"version"`
	LastKeepAlive string   `json:"lastKeepAlive"`
	Os            Os       `json:"os"`
	ID            string   `json:"id"`
	ConfigSum     string   `json:"configSum,omitempty"`
	Group         []string `json:"group,omitempty"`
	MergedSum     string   `json:"mergedSum,omitempty"`
}

func GetAllAgentsRequest(ctx context.Context, client *Client, path string) (*GetAllAgentsResponse, error) {
	response := &GetAllAgentsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetAnAgentRequest(ctx context.Context, client *Client, path string) (*GetAnAgentResponse, error) {
	response := &GetAnAgentResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-agents
func (client *Client) GetAllAgents() (*[]AgentInformation, error) {
	return client.GetAllAgentsContext(context.Background())
}

func (client *Client) GetAllAgentsContext(ctx context.Context) (*[]AgentInformation, error) {
	response, err := GetAllAgentsRequest(ctx, client, "agents")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-an-agent
func (client *Client) GetAnAgent(agentID string) (*AgentInformation, error) {
	return client.GetAnAgentContext(context.Background(), agentID)
}

func (client *Client) GetAnAgentContext(ctx context.Context, agentID string) (*AgentInformation, error) {
	response, err := GetAnAgentRequest(ctx, client, path.Join("agents", agentID))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-groups
func (client *Client) GetGroups() (*[]Group, error) {
	return client.GetGroupsContext(context.Background())
}

func (client *Client) GetGroupsContext(ctx context.Context) (*[]Group, error) {
	response, err := GetGroupsRequest(ctx, client, path.Join("agents", "groups"))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetGroupsRequest(ctx context.Context, client *Client, path string) (*GetGroupsResponse, error) {
	response := &GetGroupsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetAgentsByGroup(groupId string) (*[]AgentInformation, error) {
	return client.GetAgentsByGroupContext(context.Background(), groupId)
}

func (client *Client) GetAgentsByGroupContext(ctx context.Context, groupId string) (*[]AgentInformation, error) {
	response, err := GetAgentsByGroupRequest(ctx, client, path.Join("agents", "groups", groupId))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetAgentsByGroupRequest(ctx context.Context, client *Client, path string) (*GetAllAgentsResponse, error) {
	response := &GetAllAgentsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) RestartAgent(agentId string) (*AgentRestartData, error) {
	return client.RestartAgentContext(context.Background(), agentId)
}

func (client *Client) RestartAgentContext(ctx context.Context, agentId string) (*AgentRestartData, error) {
	response, err := RestartAgentRequest(ctx, client, path.Join("agents", agentId, "restart"))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func RestartAgentRequest(ctx context.Context, client *Client, path string) (*AgentRestartResponse, error) {
	response := &AgentRestartResponse{}
	err := PutJson(ctx, client, path, []byte{}, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
