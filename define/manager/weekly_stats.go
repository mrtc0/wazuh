package manager

type GetWeeklyStatsResponse struct {
	Error int         `json:"error"`
	Data  WeeklyStats `json:"data"`
}

type WeeklyStats struct {
	Wed struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Wed"`
	Sun struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Sun"`
	Fri struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Fri"`
	Tue struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Tue"`
	Mon struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Mon"`
	Thu struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Thu"`
	Sat struct {
		Hours        []int `json:"hours"`
		Interactions int   `json:"interactions"`
	} `json:"Sat"`
}
