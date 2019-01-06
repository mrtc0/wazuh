package define

type HardwareInfo struct {
	BoardSerial string `json:"board_serial"`
	RAM         struct {
		Usage int `json:"usage"`
		Total int `json:"total"`
		Free  int `json:"free"`
	} `json:"ram"`
	CPU struct {
		Cores int    `json:"cores"`
		Mhz   int    `json:"mhz"`
		Name  string `json:"name"`
	} `json:"cpu"`
	Scan struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
}

type GetHardwareInfoResponse struct {
	Error int          `json:"error"`
	Data  HardwareInfo `json:"data"`
}

type NetworkAddrInfo struct {
	Broadcast string `json:"broadcast,omitempty"`
	ScanID    int    `json:"scan_id"`
	Proto     string `json:"proto"`
	Netmask   string `json:"netmask"`
	Address   string `json:"address"`
}

type GetNetworkAddrInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int               `json:"totalItems"`
		Items      []NetworkAddrInfo `json:"items"`
	} `json:"data"`
}

type NetworkInterfaceInfo struct {
	Name string `json:"name"`
	Tx   struct {
		Packets int   `json:"packets"`
		Errors  int   `json:"errors"`
		Bytes   int64 `json:"bytes"`
		Dropped int   `json:"dropped"`
	} `json:"tx"`
	Scan struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
	Rx struct {
		Packets int `json:"packets"`
		Errors  int `json:"errors"`
		Bytes   int `json:"bytes"`
		Dropped int `json:"dropped"`
	} `json:"rx"`
	State string `json:"state"`
	Mtu   int    `json:"mtu"`
	Mac   string `json:"mac"`
	Type  string `json:"type"`
}

type GetNetworkInterfaceInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int                    `json:"totalItems"`
		Items      []NetworkInterfaceInfo `json:"items"`
	} `json:"data"`
}

type NetworkProtocolInfo struct {
	Dhcp    string `json:"dhcp"`
	ScanID  int    `json:"scan_id"`
	Iface   string `json:"iface"`
	Type    string `json:"type"`
	Gateway string `json:"gateway,omitempty"`
}

type GetNetworkProtocolInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int                   `json:"totalItems"`
		Items      []NetworkProtocolInfo `json:"items"`
	} `json:"data"`
}

type OSInfo struct {
	Sysname      string `json:"sysname"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Scan         struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
	Release  string `json:"release"`
	Hostname string `json:"hostname"`
	Os       struct {
		Major    string `json:"major"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
		Version  string `json:"version"`
		Codename string `json:"codename"`
		Minor    string `json:"minor"`
	} `json:"os"`
}

type GetOSInfoResponse struct {
	Error int    `json:"error"`
	Data  OSInfo `json:"data"`
}

type PackagesInfo struct {
	Vendor string `json:"vendor"`
	Name   string `json:"name"`
	Scan   struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
	Section      string `json:"section"`
	Format       string `json:"format"`
	Priority     string `json:"priority"`
	Source       string `json:"source,omitempty"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Multiarch    string `json:"multiarch,omitempty"`
	Size         int    `json:"size"`
	Description  string `json:"description"`
}

type GetPackagesInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int            `json:"totalItems"`
		Items      []PackagesInfo `json:"items"`
	} `json:"data"`
}

type PortsInfo struct {
	Remote struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"remote"`
	Scan struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
	Inode    int    `json:"inode"`
	State    string `json:"state"`
	TxQueue  int    `json:"tx_queue"`
	Protocol string `json:"protocol"`
	RxQueue  int    `json:"rx_queue"`
	Local    struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"local"`
}

type GetPortsInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int         `json:"totalItems"`
		Items      []PortsInfo `json:"items"`
	} `json:"data"`
}

type ProcessesInfo struct {
	Euser  string `json:"euser"`
	Tty    int    `json:"tty"`
	Rgroup string `json:"rgroup"`
	Sgroup string `json:"sgroup"`
	Scan   struct {
		ID   int    `json:"id"`
		Time string `json:"time"`
	} `json:"scan"`
	Resident  int    `json:"resident"`
	Share     int    `json:"share"`
	StartTime int    `json:"start_time"`
	Pid       string `json:"pid"`
	Session   int    `json:"session"`
	Stime     int    `json:"stime"`
	VMSize    int    `json:"vm_size"`
	Size      int    `json:"size"`
	Ppid      int    `json:"ppid"`
	Egroup    string `json:"egroup"`
	Name      string `json:"name"`
	Pgrp      int    `json:"pgrp"`
	Tgid      int    `json:"tgid"`
	Utime     int    `json:"utime"`
	Cmd       string `json:"cmd,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Fgroup    string `json:"fgroup"`
	State     string `json:"state"`
	Ruser     string `json:"ruser"`
	Suser     string `json:"suser"`
	Nlwp      int    `json:"nlwp"`
	Processor int    `json:"processor"`
	Nice      int    `json:"nice,omitempty"`
	Argvs     string `json:"argvs,omitempty"`
}

type GetProcessesInfoResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int             `json:"totalItems"`
		Items      []ProcessesInfo `json:"items"`
	} `json:"data"`
}
