package dtos

import "time"

type ErrorResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Code      int    `json:"code"`
}

func NewErrorResponse(message string, code int) ErrorResponse {
	return ErrorResponse{
		Message:   message,
		Timestamp: time.Now().String(),
		Code:      code,
	}
}
