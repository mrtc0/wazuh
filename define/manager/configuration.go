package manager

type GetManagerConfigurationResponse struct {
	Error int                  `json:"error"`
	Data  ManagerConfiguration `json:"data"`
}

type Alerts struct {
	LogAlertLevel   string `json:"log_alert_level"`
	EmailAlertLevel string `json:"email_alert_level"`
}

type Remote struct {
	QueueSize  string `json:"queue_size"`
	Connection string `json:"connection"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
}

type Ruleset struct {
	RuleExclude []string `json:"rule_exclude"`
	List        []string `json:"list"`
	RuleDir     []string `json:"rule_dir"`
	DecoderDir  []string `json:"decoder_dir"`
}

type CisCat struct {
	Interval    string `json:"interval"`
	Disabled    string `json:"disabled"`
	ScanOnStart string `json:"scan-on-start"`
	CiscatPath  string `json:"ciscat_path"`
	JavaPath    string `json:"java_path"`
	Timeout     string `json:"timeout"`
}

type Rootcheck struct {
	CheckUnixaudit string   `json:"check_unixaudit"`
	CheckPids      string   `json:"check_pids"`
	RootkitTrojans []string `json:"rootkit_trojans"`
	SkipNfs        string   `json:"skip_nfs"`
	CheckIf        string   `json:"check_if"`
	CheckSys       string   `json:"check_sys"`
	CheckDev       string   `json:"check_dev"`
	CheckPorts     string   `json:"check_ports"`
	Disabled       string   `json:"disabled"`
	RootkitFiles   []string `json:"rootkit_files"`
	CheckTrojans   string   `json:"check_trojans"`
	Frequency      string   `json:"frequency"`
	CheckFiles     string   `json:"check_files"`
	SystemAudit    []string `json:"system_audit"`
}

type Global struct {
	EmailNotification string   `json:"email_notification"`
	AlertsLog         string   `json:"alerts_log"`
	JsonoutOutput     string   `json:"jsonout_output"`
	SMTPServer        string   `json:"smtp_server"`
	QueueSize         string   `json:"queue_size"`
	EmailTo           string   `json:"email_to"`
	Logall            string   `json:"logall"`
	EmailMaxperhour   string   `json:"email_maxperhour"`
	WhiteList         []string `json:"white_list"`
	EmailFrom         string   `json:"email_from"`
	LogallJSON        string   `json:"logall_json"`
}

type Integration struct {
	AlertFormat string `json:"alert_format"`
	HookURL     string `json:"hook_url"`
	Name        string `json:"name"`
	Level       string `json:"level"`
}

type Auth struct {
	Purge            string `json:"purge"`
	Ciphers          string `json:"ciphers"`
	ForceInsert      string `json:"force_insert"`
	SslVerifyHost    string `json:"ssl_verify_host"`
	LimitMaxagents   string `json:"limit_maxagents"`
	ForceTime        string `json:"force_time"`
	SslManagerKey    string `json:"ssl_manager_key"`
	Disabled         string `json:"disabled"`
	SslManagerCert   string `json:"ssl_manager_cert"`
	UseSourceIP      string `json:"use_source_ip"`
	UsePassword      string `json:"use_password"`
	Port             string `json:"port"`
	SslAutoNegotiate string `json:"ssl_auto_negotiate"`
}

type Syscollector struct {
	Hardware    string `json:"hardware"`
	Processes   string `json:"processes"`
	Network     string `json:"network"`
	Interval    string `json:"interval"`
	ScanOnStart string `json:"scan_on_start"`
	Disabled    string `json:"disabled"`
	Packages    string `json:"packages"`
	Os          string `json:"os"`
	Ports       Ports  `json:"ports"`
}

type Ports struct {
	Item string `json:"item"`
	All  string `json:"all"`
}

type Syscheck struct {
	Ignore        []string      `json:"ignore"`
	SkipNfs       string        `json:"skip_nfs"`
	Directories   []Directories `json:"directories"`
	ScanOnStart   string        `json:"scan_on_start"`
	AlertNewFiles string        `json:"alert_new_files"`
	Disabled      string        `json:"disabled"`
	Frequency     string        `json:"frequency"`
	RestartAudit  string        `json:"restart_audit"`
	AutoIgnore    AutoIgnore    `json:"auto_ignore"`
	RemoveOldDiff string        `json:"remove_old_diff"`
	Nodiff        []string      `json:"nodiff"`
}

type Directories struct {
	Path     string `json:"path"`
	CheckAll string `json:"check_all"`
}

type AutoIgnore struct {
	Item      string `json:"item"`
	Frequency string `json:"frequency"`
	Timeframe string `json:"timeframe"`
}

type VulnerabilityDetector struct {
	Disabled   string `json:"disabled"`
	IgnoreTime string `json:"ignore_time"`
	Feed       []Feed `json:"feed"`
	Interval   string `json:"interval"`
	RunOnStart string `json:"run_on_start"`
}

type Feed struct {
	Disabled       string `json:"disabled"`
	UpdateInterval string `json:"update_interval"`
	Name           string `json:"name"`
}

type OpenScap struct {
	Disabled    string `json:"disabled"`
	ScanOnStart string `json:"scan-on-start"`
	Interval    string `json:"interval"`
	Timeout     string `json:"timeout"`
}

type Cluster struct {
	Disabled string   `json:"disabled"`
	Hidden   string   `json:"hidden"`
	Name     string   `json:"name"`
	NodeName string   `json:"node_name"`
	BindAddr string   `json:"bind_addr"`
	NodeType string   `json:"node_type"`
	Nodes    []string `json:"nodes"`
	Port     string   `json:"port"`
}

type Command struct {
	Executable     string `json:"executable"`
	TimeoutAllowed string `json:"timeout_allowed,omitempty"`
	Name           string `json:"name"`
	Expect         string `json:"expect,omitempty"`
}

type Localfile struct {
	LogFormat string `json:"log_format"`
	Frequency string `json:"frequency,omitempty"`
	Command   string `json:"command,omitempty"`
	Alias     string `json:"alias,omitempty"`
	Location  string `json:"location,omitempty"`
}

type Osquery struct {
	Disabled   string `json:"disabled"`
	LogPath    string `json:"log_path"`
	AddLabels  string `json:"add_labels"`
	RunDaemon  string `json:"run_daemon"`
	ConfigPath string `json:"config_path"`
}

type ManagerConfiguration struct {
	Alerts                Alerts                `json:"alerts"`
	Remote                []Remote              `json:"remote"`
	Ruleset               Ruleset               `json:"ruleset"`
	CisCat                CisCat                `json:"cis-cat"`
	Rootcheck             Rootcheck             `json:"rootcheck"`
	Global                Global                `json:"global"`
	Integration           []Integration         `json:"integration"`
	Auth                  Auth                  `json:"auth"`
	Syscollector          Syscollector          `json:"syscollector"`
	Syscheck              Syscheck              `json:"syscheck"`
	VulnerabilityDetector VulnerabilityDetector `json:"vulnerability-detector"`
	OpenScap              OpenScap              `json:"open-scap"`
	Cluster               Cluster               `json:"cluster"`
	Command               []Command             `json:"command"`
	Localfile             []Localfile           `json:"localfile"`
	Osquery               Osquery               `json:"osquery"`
}
