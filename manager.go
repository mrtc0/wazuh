package wazuh

import (
	"context"

	. "github.com/mrtc0/wazuh/define/manager"
)

func GetManagerConfigurationRequest(ctx context.Context, client httpClient, path string) (*GetManagerConfigurationResponse, error) {
	response := &GetManagerConfigurationResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-configuration
func (api *Client) GetManagerConfiguration() (*ManagerConfiguration, error) {
	return api.GetManagerConfigurationContext(context.Background())

}

func (api *Client) GetManagerConfigurationContext(ctx context.Context) (*ManagerConfiguration, error) {
	response, err := GetManagerConfigurationRequest(ctx, api.httpclient, "manager/configuration")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerInformationRequest(ctx context.Context, client httpClient, path string) (*GetManagerInformationResponse, error) {
	response := &GetManagerInformationResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-information
func (api *Client) GetManagerInformation() (*ManagerInformation, error) {
	return api.GetManagerInformationContext(context.Background())

}

func (api *Client) GetManagerInformationContext(ctx context.Context) (*ManagerInformation, error) {
	response, err := GetManagerInformationRequest(ctx, api.httpclient, "manager/info")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetLogsRequest(ctx context.Context, client httpClient, path string) (*GetLogsResponse, error) {
	response := &GetLogsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-ossec-log
func (api *Client) GetLogs() (*[]Log, error) {
	return api.GetLogsContext(context.Background())
}

func (api *Client) GetLogsContext(ctx context.Context) (*[]Log, error) {
	response, err := GetLogsRequest(ctx, api.httpclient, "manager/logs")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetLogSummaryRequest(ctx context.Context, client httpClient, path string) (*GetLogSummaryResponse, error) {
	response := &GetLogSummaryResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-summary-of-ossec-log
func (api *Client) GetLogSummary() (*LogSummary, error) {
	return api.GetLogSummaryContext(context.Background())

}

func (api *Client) GetLogSummaryContext(ctx context.Context) (*LogSummary, error) {
	response, err := GetLogSummaryRequest(ctx, api.httpclient, "manager/logs/summary")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerStatsRequest(ctx context.Context, client httpClient, path string) (*GetManagerStatsResponse, error) {
	response := &GetManagerStatsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats
func (api *Client) GetManagerStats() (*[]ManagerStats, error) {
	return api.GetManagerStatsContext(context.Background())

}

func (api *Client) GetManagerStatsContext(ctx context.Context) (*[]ManagerStats, error) {
	response, err := GetManagerStatsRequest(ctx, api.httpclient, "manager/stats")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetAnalysisdStatsRequest(ctx context.Context, client httpClient, path string) (*GetAnalysisdStatsResponse, error) {
	response := &GetAnalysisdStatsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-analysisd-stats
func (api *Client) GetAnalysisdStats() (*AnalysisdStats, error) {
	return api.GetAnalysisdStatsContext(context.Background())
}

func (api *Client) GetAnalysisdStatsContext(ctx context.Context) (*AnalysisdStats, error) {
	response, err := GetAnalysisdStatsRequest(ctx, api.httpclient, "manager/stats/analysisd")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetStatsByHourRequest(ctx context.Context, client httpClient, path string) (*GetStatsByHourResponse, error) {
	response := &GetStatsByHourResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-hour

func (api *Client) GetStatsByHour() (*StatsByHour, error) {
	return api.GetStatsByHourContext(context.Background())
}

func (api *Client) GetStatsByHourContext(ctx context.Context) (*StatsByHour, error) {
	response, err := GetStatsByHourRequest(ctx, api.httpclient, "manager/stats/hourly")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetRemotedStatsRequest(ctx context.Context, client httpClient, path string) (*GetRemotedStatsResponse, error) {
	response := &GetRemotedStatsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-week
func (api *Client) GetRemotedStats() (*RemotedStats, error) {
	return api.GetRemotedStatsContext(context.Background())
}

func (api *Client) GetRemotedStatsContext(ctx context.Context) (*RemotedStats, error) {
	response, err := GetRemotedStatsRequest(ctx, api.httpclient, "manager/stats/remoted")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetWeeklyStatsRequest(ctx context.Context, client httpClient, path string) (*GetWeeklyStatsResponse, error) {
	response := &GetWeeklyStatsResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err

	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-week
func (api *Client) GetWeeklyStats() (*WeeklyStats, error) {
	return api.GetWeeklyStatsContext(context.Background())
}

func (api *Client) GetWeeklyStatsContext(ctx context.Context) (*WeeklyStats, error) {
	response, err := GetWeeklyStatsRequest(ctx, api.httpclient, "manager/stats/weekly")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerStatusRequest(ctx context.Context, client httpClient, path string) (*GetManagerStatusResponse, error) {
	response := &GetManagerStatusResponse{}
	err := GetJson(ctx, client, APIURL+path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-status
func (api *Client) GetManagerStatus() (*ManagerStatus, error) {
	return api.GetManagerStatusContext(context.Background())
}

func (api *Client) GetManagerStatusContext(ctx context.Context) (*ManagerStatus, error) {
	response, err := GetManagerStatusRequest(ctx, api.httpclient, "manager/status")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}
