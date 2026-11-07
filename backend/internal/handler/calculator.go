package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/sezzle/calculator-backend/internal/service"
)

type CalculatorHandler struct {
	service   *service.CalculatorService
	validator *validator.Validate
	logger    *slog.Logger
}

func NewCalculatorHandler(svc *service.CalculatorService, v *validator.Validate, logger *slog.Logger) *CalculatorHandler {
	return &CalculatorHandler{service: svc, validator: v, logger: logger}
}

// Calculate godoc
// @Summary Perform a calculation
// @Description Execute an arithmetic operation with one or two operands
// @Tags Calculator
// @Accept json
// @Produce json
// @Param request body dto.CalculateRequest true "Calculation request"
// @Success 200 {object} dto.CalculateResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Router /api/calculate [post]
func (h *CalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req dto.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusUnprocessableEntity, "INVALID_INPUT", "Invalid JSON payload", nil)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		h.writeError(w, http.StatusUnprocessableEntity, "INVALID_INPUT", "Validation failed: "+err.Error(), nil)
		return
	}
	resp, err := h.service.Calculate(req)
	if err != nil {
		h.handleServiceError(w, err, req)
		return
	}
	h.writeJSON(w, http.StatusOK, resp)
}

func (h *CalculatorHandler) handleServiceError(w http.ResponseWriter, err error, req dto.CalculateRequest) {
	details := map[string]interface{}{"operation": req.Operation, "operand_a": req.OperandA}
	if req.OperandB != nil {
		details["operand_b"] = *req.OperandB
	}
	var domErr *domain.DomainError
	if errors.As(err, &domErr) {
		status := http.StatusUnprocessableEntity
		if domErr.Code == "INVALID_OPERATION" {
			status = http.StatusBadRequest
		}
		h.writeError(w, status, domErr.Code, domErr.Message, details)
		return
	}
	h.logger.Error("unexpected error", "error", err.Error())
	h.writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred", nil)
}

func (h *CalculatorHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *CalculatorHandler) writeError(w http.ResponseWriter, status int, code, message string, details map[string]interface{}) {
	h.writeJSON(w, status, dto.NewErrorResponse(code, message, details))
}
