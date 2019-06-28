package wazuh

import (
	"context"
	"path"

	. "github.com/mrtc0/wazuh/define"
)

func GetRulesRequest(ctx context.Context, client *Client, path string) (*GetRulesResponse, error) {
	response := &GetRulesResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetRuleFilesRequest(ctx context.Context, client *Client, path string) (*GetRuleFilesResponse, error) {
	response := &GetRuleFilesResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetRequest(ctx context.Context, client *Client, path string) (*Response, error) {
	response := &Response{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-all-rules
func (client *Client) GetAllRules() (*[]Rules, error) {
	return client.GetAllRulesContext(context.Background())
}

func (client *Client) GetAllRulesContext(ctx context.Context) (*[]Rules, error) {
	response, err := GetRulesRequest(ctx, client, "rules")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rules-by-id
func (client *Client) GetRulesById(ruleId string) (*[]Rules, error) {
	return client.GetRulesByIdContext(context.Background(), ruleId)
}

func (client *Client) GetRulesByIdContext(ctx context.Context, ruleId string) (*[]Rules, error) {
	response, err := GetRulesRequest(ctx, client, path.Join("rules", ruleId))
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-files-of-rules
func (client *Client) GetRuleFiles() (*[]RuleFiles, error) {
	return client.GetRuleFilesContext(context.Background())
}

func (client *Client) GetRuleFilesContext(ctx context.Context) (*[]RuleFiles, error) {
	response, err := GetRuleFilesRequest(ctx, client, "rules/files")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-gdpr-requirements
func (client *Client) GetGdprRules() (*[]string, error) {
	return client.GetGdprRulesContext(context.Background())
}

func (client *Client) GetGdprRulesContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, client, "rules/gdpr")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-groups
func (client *Client) GetRuleGroups() (*[]string, error) {
	return client.GetRuleGroupsContext(context.Background())
}

func (client *Client) GetRuleGroupsContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, client, "rules/groups")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-rule-pci-requirements
func (client *Client) GetRulePCI() (*[]string, error) {
	return client.GetRulePCIContext(context.Background())
}

func (client *Client) GetRulePCIContext(ctx context.Context) (*[]string, error) {
	response, err := GetRequest(ctx, client, "rules/pci")
	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}
