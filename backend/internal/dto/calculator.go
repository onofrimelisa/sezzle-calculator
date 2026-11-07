package dto

import "time"

type CalculateRequest struct {
	Operation string   `json:"operation" validate:"required"`
	OperandA  float64  `json:"operand_a" validate:"required"`
	OperandB  *float64 `json:"operand_b,omitempty"`
}

type CalculateResponse struct {
	Operation string   `json:"operation"`
	OperandA  float64  `json:"operand_a"`
	OperandB  *float64 `json:"operand_b,omitempty"`
	Result    float64  `json:"result"`
	Timestamp string   `json:"timestamp"`
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
