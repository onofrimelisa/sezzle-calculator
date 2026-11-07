package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/sezzle/calculator-backend/docs"
	"github.com/sezzle/calculator-backend/internal/domain"
	"github.com/sezzle/calculator-backend/internal/handler"
	"github.com/sezzle/calculator-backend/internal/middleware"
	"github.com/sezzle/calculator-backend/internal/operations"
	"github.com/sezzle/calculator-backend/internal/platform/config"
	"github.com/sezzle/calculator-backend/internal/platform/logger"
	"github.com/sezzle/calculator-backend/internal/platform/metrics"
	"github.com/sezzle/calculator-backend/internal/service"
)

// @title Calculator API
// @version 1.0
// @description A RESTful calculator API with basic and advanced arithmetic operations.
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	m := metrics.New()

	registry := buildRegistry()
	svc := service.NewCalculatorService(registry, log, m)
	v := validator.New()
	calcHandler := handler.NewCalculatorHandler(svc, v, log)

	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.Recovery(log))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logging(log))
	r.Use(middleware.Metrics(m))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.CORSOrigins},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	r.Post("/api/calculate", calcHandler.Calculate)
	r.Get("/healthz", handler.HealthCheck)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		log.Info("server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server failed", "error", err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("server shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced shutdown", "error", err.Error())
		os.Exit(1)
	}
	log.Info("server stopped")
}

func buildRegistry() *domain.Registry {
	registry := domain.NewRegistry()
	registry.RegisterBinary(operations.NewAdd())
	registry.RegisterBinary(operations.NewSubtract())
	registry.RegisterBinary(operations.NewMultiply())
	registry.RegisterBinary(operations.NewDivide())
	registry.RegisterBinary(operations.NewPower())
	registry.RegisterUnary(operations.NewSqrt())
	registry.RegisterBinary(operations.NewPercentage())
	return registry
}
