package model

type Return struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
