package dto

import "time"

type CalculateRequest struct {
	Operation string   `json:"operation" validate:"required" example:"add"`
	OperandA  float64  `json:"operand_a" validate:"required" example:"5"`
	OperandB  *float64 `json:"operand_b,omitempty" example:"3"`
}

type CalculateResponse struct {
	Operation string   `json:"operation" example:"add"`
	OperandA  float64  `json:"operand_a" example:"5"`
	OperandB  *float64 `json:"operand_b,omitempty" example:"3"`
	Result    float64  `json:"result" example:"8"`
	Timestamp string   `json:"timestamp" example:"2026-11-07T15:04:05Z"`
}

func NewCalculateResponse(operation string, a float64, b *float64, result float64) CalculateResponse {
	return CalculateResponse{
		Operation: operation,
		OperandA:  a,
		OperandB:  b,
		Result:    result,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

type ErrorDetail struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func NewErrorResponse(code, message string, details map[string]interface{}) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
}
