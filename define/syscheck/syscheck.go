package syscheck

type GetSyscheckFilesResponse struct {
	Data struct {
		Items      []SyscheckFiles `json:"items"`
		TotalItems int64           `json:"totalItems"`
	} `json:"data"`
	Error int64 `json:"error"`
}

type SyscheckFiles struct {
	Attributes int64  `json:"attributes"`
	Date       string `json:"date"`
	File       string `json:"file"`
	Gid        string `json:"gid"`
	Gname      string `json:"gname"`
	Inode      int64  `json:"inode"`
	Md5        string `json:"md5"`
	Mtime      string `json:"mtime"`
	Perm       string `json:"perm"`
	Sha1       string `json:"sha1"`
	Sha256     string `json:"sha256"`
	Size       int64  `json:"size"`
	Type       string `json:"type"`
	UID        string `json:"uid"`
	Uname      string `json:"uname"`
}

type RunSyscheckAllAgentsResponse struct {
	Data  string `json:"data"`
	Error int64  `json:"error"`
}

type RunSyscheckAgentResponse struct {
	Data  string `json:"data"`
	Error int64  `json:"error"`
}

type GetLastSyscheckScanResponse struct {
	Data  LastSyscheckScan `json:"data"`
	Error int64            `json:"error"`
}

type LastSyscheckScan struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type ClearSyscheckDatabaseResponse struct {
	Data  string `json:"data"`
	Error int64  `json:"error"`
}
