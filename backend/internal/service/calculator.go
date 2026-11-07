package service

import (
	"fmt"
	"log/slog"

	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/dto"
	"github.com/sezzle/calculator-backend/internal/platform/metrics"
)

type CalculatorService struct {
	registry *domain.Registry
	logger   *slog.Logger
	metrics  *metrics.Metrics
}

func NewCalculatorService(registry *domain.Registry, logger *slog.Logger, m *metrics.Metrics) *CalculatorService {
	return &CalculatorService{registry: registry, logger: logger, metrics: m}
}

func (s *CalculatorService) Calculate(req dto.CalculateRequest) (dto.CalculateResponse, error) {
	entry, err := s.registry.Get(req.Operation)
	if err != nil {
		s.logger.Warn("invalid operation requested", "operation", req.Operation)
		return dto.CalculateResponse{}, err
	}

	var result float64
	switch entry.Type {
	case domain.Binary:
		if req.OperandB == nil {
			return dto.CalculateResponse{}, domain.ErrMissingOperand
		}
		result, err = entry.Binary.Execute(req.OperandA, *req.OperandB)
	case domain.Unary:
		if req.OperandB != nil {
			return dto.CalculateResponse{}, domain.ErrUnexpectedOperand
		}
		result, err = entry.Unary.Execute(req.OperandA)
	default:
		return dto.CalculateResponse{}, fmt.Errorf("unknown operation type: %s", entry.Type)
	}

	if err != nil {
		s.logger.Warn("operation failed", "operation", req.Operation, "operand_a", req.OperandA, "error", err.Error())
		return dto.CalculateResponse{}, err
	}

	s.logger.Debug("calculation successful", "operation", req.Operation, "operand_a", req.OperandA, "result", result)
	return dto.NewCalculateResponse(req.Operation, req.OperandA, req.OperandB, result), nil
}
