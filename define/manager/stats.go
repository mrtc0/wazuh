package manager

type GetManagerStatsResponse struct {
	Error int            `json:"error"`
	Data  []ManagerStats `json:"data"`
}

type ManagerStats struct {
	Hour     int `json:"hour"`
	Firewall int `json:"firewall"`
	Alerts   []struct {
		Times int `json:"times"`
		Sigid int `json:"sigid"`
		Level int `json:"level"`
	} `json:"alerts"`
	TotalAlerts int `json:"totalAlerts"`
	Syscheck    int `json:"syscheck"`
	Events      int `json:"events"`
}
