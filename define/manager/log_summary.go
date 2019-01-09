package manager

type GetLogSummaryResponse struct {
	Error int        `json:"error"`
	Data  LogSummary `json:"data"`
}

type OssecSyscheckd struct {
	Info     int `json:"info"`
	All      int `json:"all"`
	Critical int `json:"critical"`
	Error    int `json:"error"`
	Debug    int `json:"debug"`
	Warning  int `json:"warning"`
}

type WazuhModulesdSyscollector struct {
	Info     int `json:"info"`
	All      int `json:"all"`
	Critical int `json:"critical"`
	Error    int `json:"error"`
	Debug    int `json:"debug"`
	Warning  int `json:"warning"`
}

type OssecRootcheck struct {
	Info     int `json:"info"`
	All      int `json:"all"`
	Critical int `json:"critical"`
	Error    int `json:"error"`
	Debug    int `json:"debug"`
	Warning  int `json:"warning"`
}

type OssecMonitord struct {
	Info     int `json:"info"`
	All      int `json:"all"`
	Critical int `json:"critical"`
	Error    int `json:"error"`
	Debug    int `json:"debug"`
	Warning  int `json:"warning"`
}

type LogSummary struct {
	OssecSyscheckd            OssecSyscheckd            `json:"ossec-syscheckd"`
	WazuhModulesdSyscollector WazuhModulesdSyscollector `json:"wazuh-modulesd:syscollector"`
	OssecRootcheck            OssecRootcheck            `json:"ossec-rootcheck"`
	OssecMonitord             OssecMonitord             `json:"ossec-monitord"`
}
