package wazuh

import (
	"context"
	"strconv"

	. "github.com/mrtc0/wazuh/define"
)

func GetRulesRequest(ctx context.Context, client httpClient, path string) (*GetRulesResponse, error) {
	response := &GetRulesResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetRuleFilesRequest(ctx context.Context, client httpClient, path string) (*GetRuleFilesResponse, error) {
	response := &GetRuleFilesResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetRequest(ctx context.Context, client httpClient, path string) (*Response, error) {
	response := &Response{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-rules
func (api *Client) GetAllRules() (*[]Rules, error) {
	return api.GetAllRulesContext(context.Background())
}

func (api *Client) GetAllRulesContext(ctx context.Context) (*[]Rules, error) {
	response, err := GetRulesRequest(ctx, api.httpclient, "rules")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rules-by-id
func (api *Client) GetRulesById(ruleId int) (*[]Rules, error) {
	return api.GetRulesByIdContext(context.Background(), ruleId)
}

func (api *Client) GetRulesByIdContext(ctx context.Context, ruleId int) (*[]Rules, error) {
	response, err := GetRulesRequest(ctx, api.httpclient, "rules/"+strconv.Itoa(ruleId))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-files-of-rules
func (api *Client) GetRuleFiles() (*[]RuleFiles, error) {
	return api.GetRuleFilesContext(context.Background())
}

func (api *Client) GetRuleFilesContext(ctx context.Context) (*[]RuleFiles, error) {
	response, err := GetRuleFilesRequest(ctx, api.httpclient, "rules/files")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-gdpr-requirements
func (api *Client) GetGdprRules() (*[]string, error) {
	return api.GetGdprRulesContext(context.Background())
}

func (api *Client) GetGdprRulesContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, api.httpclient, "rules/gdpr")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-groups
func (api *Client) GetRuleGroups() (*[]string, error) {
	return api.GetRuleGroupsContext(context.Background())
}

func (api *Client) GetRuleGroupsContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, api.httpclient, "rules/groups")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-pci-requirements
func (api *Client) GetRulePCI() (*[]string, error) {
	return api.GetRulePCIContext(context.Background())
}

func (api *Client) GetRulePCIContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, api.httpclient, "rules/pci")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
