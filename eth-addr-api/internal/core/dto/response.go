package dto

type BaseOKResponse struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type BaseOKResponseWithData[T any] struct {
	BaseOKResponse
	Data T `json:"data,omitempty"`
}