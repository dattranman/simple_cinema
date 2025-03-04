package response

const (
	CodeSuccess int = iota + 1
	CodeError
	CodeValidationError
)

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type BaseList struct {
	Base
	Total int `json:"total"`
}
