package define

type Rules struct {
	Status      string        `json:"status"`
	Pci         []interface{} `json:"pci"`
	Description string        `json:"description"`
	File        string        `json:"file"`
	Level       int           `json:"level"`
	Path        string        `json:"path"`
	Details     struct {
		Category string `json:"category"`
		Noalert  string `json:"noalert"`
	} `json:"details"`
	Groups []string      `json:"groups"`
	ID     int           `json:"id"`
	Gdpr   []interface{} `json:"gdpr"`
}

type GetRulesResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int     `json:"totalItems"`
		Items      []Rules `json:"items"`
	} `json:"data"`
}

type RuleFiles struct {
	Status string `json:"status"`
	Path   string `json:"path"`
	File   string `json:"file"`
}

type GetRuleFilesResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int         `json:"totalItems"`
		Items      []RuleFiles `json:"items"`
	} `json:"data"`
}

// GDPR, Groups, PCI
type Response struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int      `json:"totalItems"`
		Items      []string `json:"items"`
	} `json:"data"`
}
