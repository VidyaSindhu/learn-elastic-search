package dtos

type BaseResponseDto struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type ErrorResponse struct {
	BaseResponseDto
	ErrorCode int `json:"error_code,omitempty"`
}

type SuccessResponseDto struct {
	BaseResponseDto
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
