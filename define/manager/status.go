package manager

type GetManagerStatusResponse struct {
	Error int           `json:"error"`
	Data  ManagerStatus `json:"data"`
}

type ManagerStatus struct {
	WazuhModulesd     string `json:"wazuh-modulesd"`
	OssecAuthd        string `json:"ossec-authd"`
	WazuhClusterd     string `json:"wazuh-clusterd"`
	OssecMonitord     string `json:"ossec-monitord"`
	OssecLogcollector string `json:"ossec-logcollector"`
	OssecExecd        string `json:"ossec-execd"`
	OssecRemoted      string `json:"ossec-remoted"`
	OssecSyscheckd    string `json:"ossec-syscheckd"`
	OssecAnalysisd    string `json:"ossec-analysisd"`
	OssecMaild        string `json:"ossec-maild"`
}
