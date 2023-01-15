package tcp_model

type Request struct {
	Category string      `json:"category"`
	Method   string      `json:"method"`
	Args     interface{} `json:"args"`
}
