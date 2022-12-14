package types

type ApiResponse struct {
	Status  int         `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
