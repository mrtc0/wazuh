package manager

type GetStatsByHourResponse struct {
	Error int         `json:"error"`
	Data  StatsByHour `json:"data"`
}

type StatsByHour struct {
	Averages     []int `json:"averages"`
	Interactions int   `json:"interactions"`
}
