package model

type Return struct {
	Data   interface{} `json:"data"`
	Error  Error       `json:"error"`
	Status string      `json:"status"`
}

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
