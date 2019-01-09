package manager

type ManagerInformation struct {
	CompilationDate string `json:"compilation_date"`
	Version         string `json:"version"`
	OpensslSupport  string `json:"openssl_support"`
	MaxAgents       string `json:"max_agents"`
	RulesetVersion  string `json:"ruleset_version"`
	Path            string `json:"path"`
	TzName          string `json:"tz_name"`
	Type            string `json:"type"`
	TzOffset        string `json:"tz_offset"`
}

type GetManagerInformationResponse struct {
	Error int                `json:"error"`
	Data  ManagerInformation `json:"data"`
}
