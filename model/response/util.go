package response

const (
	CodeSuccess int = iota + 1
	CodeError
	CodeValidationError
)

type Base struct {
	Code    int    `json:"code" example:"1"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

type BaseList struct {
	Base
	Total int `json:"total"`
}
