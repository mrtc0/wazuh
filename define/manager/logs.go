package manager

type Log struct {
	Timestamp   string `json:"timestamp"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

type GetLogsResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int   `json:"totalItems"`
		Items      []Log `json:"items"`
	} `json:"data"`
}
