package sca

type GetScaChecksResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int        `json:"totalItems"`
		Items      []ScaCheck `json:"items"`
	} `json:"data"`
}

type ScaCheck struct {
	Rationale   string `json:"rationale"`
	Remediation string `json:"remediation"`
	File        string `json:"file"`
	Title       string `json:"title"`
	Result      string `json:"result"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	PolicyID    string `json:"policy_id"`
	Rules       []struct {
		Type string `json:"type"`
		Rule string `json:"rule"`
	} `json:"rules"`
}

type GetScaResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int   `json:"totalItems"`
		Items      []Sca `json:"items"`
	} `json:"data"`
}

type Sca struct {
	Score       int    `json:"score"`
	EndScan     string `json:"end_scan"`
	HashFile    string `json:"hash_file"`
	Pass        int    `json:"pass"`
	Fail        int    `json:"fail"`
	References  string `json:"references"`
	StartScan   string `json:"start_scan"`
	PolicyID    string `json:"policy_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Invalid     int    `json:"invalid"`
	TotalChecks int    `json:"total_checks"`
}
