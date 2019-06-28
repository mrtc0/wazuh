package wazuh

import (
	"context"

	. "github.com/mrtc0/wazuh/define/manager"
)

func GetManagerConfigurationRequest(ctx context.Context, client *Client, path string) (*GetManagerConfigurationResponse, error) {
	response := &GetManagerConfigurationResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-configuration
func (client *Client) GetManagerConfiguration() (*ManagerConfiguration, error) {
	return client.GetManagerConfigurationContext(context.Background())

}

func (client *Client) GetManagerConfigurationContext(ctx context.Context) (*ManagerConfiguration, error) {
	response, err := GetManagerConfigurationRequest(ctx, client, "manager/configuration")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerInformationRequest(ctx context.Context, client *Client, path string) (*GetManagerInformationResponse, error) {
	response := &GetManagerInformationResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-information
func (client *Client) GetManagerInformation() (*ManagerInformation, error) {
	return client.GetManagerInformationContext(context.Background())

}

func (client *Client) GetManagerInformationContext(ctx context.Context) (*ManagerInformation, error) {
	response, err := GetManagerInformationRequest(ctx, client, "manager/info")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetLogsRequest(ctx context.Context, client *Client, path string) (*GetLogsResponse, error) {
	response := &GetLogsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-ossec-log
func (client *Client) GetLogs() (*[]Log, error) {
	return client.GetLogsContext(context.Background())
}

func (client *Client) GetLogsContext(ctx context.Context) (*[]Log, error) {
	response, err := GetLogsRequest(ctx, client, "manager/logs")

	if err != nil {
		return nil, err
	}
	return &response.Data.Items, nil
}

func GetLogSummaryRequest(ctx context.Context, client *Client, path string) (*GetLogSummaryResponse, error) {
	response := &GetLogSummaryResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-summary-of-ossec-log
func (client *Client) GetLogSummary() (*LogSummary, error) {
	return client.GetLogSummaryContext(context.Background())

}

func (client *Client) GetLogSummaryContext(ctx context.Context) (*LogSummary, error) {
	response, err := GetLogSummaryRequest(ctx, client, "manager/logs/summary")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerStatsRequest(ctx context.Context, client *Client, path string) (*GetManagerStatsResponse, error) {
	response := &GetManagerStatsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats
func (client *Client) GetManagerStats() (*[]ManagerStats, error) {
	return client.GetManagerStatsContext(context.Background())

}

func (client *Client) GetManagerStatsContext(ctx context.Context) (*[]ManagerStats, error) {
	response, err := GetManagerStatsRequest(ctx, client, "manager/stats")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetAnalysisdStatsRequest(ctx context.Context, client *Client, path string) (*GetAnalysisdStatsResponse, error) {
	response := &GetAnalysisdStatsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-analysisd-stats
func (client *Client) GetAnalysisdStats() (*AnalysisdStats, error) {
	return client.GetAnalysisdStatsContext(context.Background())
}

func (client *Client) GetAnalysisdStatsContext(ctx context.Context) (*AnalysisdStats, error) {
	response, err := GetAnalysisdStatsRequest(ctx, client, "manager/stats/analysisd")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetStatsByHourRequest(ctx context.Context, client *Client, path string) (*GetStatsByHourResponse, error) {
	response := &GetStatsByHourResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-hour

func (client *Client) GetStatsByHour() (*StatsByHour, error) {
	return client.GetStatsByHourContext(context.Background())
}

func (client *Client) GetStatsByHourContext(ctx context.Context) (*StatsByHour, error) {
	response, err := GetStatsByHourRequest(ctx, client, "manager/stats/hourly")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetRemotedStatsRequest(ctx context.Context, client *Client, path string) (*GetRemotedStatsResponse, error) {
	response := &GetRemotedStatsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-week
func (client *Client) GetRemotedStats() (*RemotedStats, error) {
	return client.GetRemotedStatsContext(context.Background())
}

func (client *Client) GetRemotedStatsContext(ctx context.Context) (*RemotedStats, error) {
	response, err := GetRemotedStatsRequest(ctx, client, "manager/stats/remoted")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetWeeklyStatsRequest(ctx context.Context, client *Client, path string) (*GetWeeklyStatsResponse, error) {
	response := &GetWeeklyStatsResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err

	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-stats-by-week
func (client *Client) GetWeeklyStats() (*WeeklyStats, error) {
	return client.GetWeeklyStatsContext(context.Background())
}

func (client *Client) GetWeeklyStatsContext(ctx context.Context) (*WeeklyStats, error) {
	response, err := GetWeeklyStatsRequest(ctx, client, "manager/stats/weekly")

	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func GetManagerStatusRequest(ctx context.Context, client *Client, path string) (*GetManagerStatusResponse, error) {
	response := &GetManagerStatusResponse{}
	err := GetJson(ctx, client, path, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// https://documentation.wazuh.com/current/user-manual/api/reference.html#get-manager-status
func (client *Client) GetManagerStatus() (*ManagerStatus, error) {
	return client.GetManagerStatusContext(context.Background())
}

func (client *Client) GetManagerStatusContext(ctx context.Context) (*ManagerStatus, error) {
	response, err := GetManagerStatusRequest(ctx, client, "manager/status")
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}
