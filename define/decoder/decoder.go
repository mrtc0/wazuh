package decoder

type DecoderDetails interface{}

type Decoder struct {
	Status   string         `json:"status"`
	Name     string         `json:"name"`
	Details  DecoderDetails `json:"details"`
	File     string         `json:"file"`
	Position int            `json:"position"`
	Path     string         `json:"path"`
}

type GetDecodersResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int       `json:"totalItems"`
		Items      []Decoder `json:"items"`
	} `json:"data"`
}

type DecoderFiles struct {
	Status string `json:"status"`
	Path   string `json:"path"`
	File   string `json:"file"`
}

type GetDecoderFilesResponse struct {
	Error int `json:"error"`
	Data  struct {
		TotalItems int            `json:"totalItems"`
		Items      []DecoderFiles `json:"items"`
	} `json:"data"`
}
