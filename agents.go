package wazuh

import (
	"context"
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

func GetAllAgentsRequest(ctx context.Context, client httpClient, path string) (*GetAllAgentsResponse, error) {
	response := &GetAllAgentsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetAnAgentRequest(ctx context.Context, client httpClient, path string) (*GetAnAgentResponse, error) {
	response := &GetAnAgentResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-agents
func (api *Client) GetAllAgents() (*[]AgentInformation, error) {
	return api.GetAllAgentsContext(context.Background())
}

func (api *Client) GetAllAgentsContext(ctx context.Context) (*[]AgentInformation, error) {
	response, err := GetAllAgentsRequest(ctx, api.httpclient, "agents")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-an-agent
func (api *Client) GetAnAgent(agentID string) (*AgentInformation, error) {
	return api.GetAnAgentContext(context.Background(), agentID)
}

func (api *Client) GetAnAgentContext(ctx context.Context, agentID string) (*AgentInformation, error) {
	response, err := GetAnAgentRequest(ctx, api.httpclient, "agents/"+agentID)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-groups
func (api *Client) GetGroups() (*[]Group, error) {
	return api.GetGroupsContext(context.Background())
}

func (api *Client) GetGroupsContext(ctx context.Context) (*[]Group, error) {
	response, err := GetGroupsRequest(ctx, api.httpclient, "agents/groups")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetGroupsRequest(ctx context.Context, client httpClient, path string) (*GetGroupsResponse, error) {
	response := &GetGroupsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
