package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

const (
	StatusOk    = "ok"
	StatusError = "error"
)
