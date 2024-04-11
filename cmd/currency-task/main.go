package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/olegtemek/currency-task/internal/app"
	"github.com/olegtemek/currency-task/internal/config"
)

// @title CurrencyService
// @version 1.0

// @Host localhost:8000
// @BasePath /
func main() {
	ctx := context.Background()
	cfg := config.New()
	log := setUpLogger(cfg.Env)

	application := app.New(log, cfg)
	server := application.Init()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("ERROR", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	if err := server.Shutdown(ctx); err != nil {
		log.Info("Could not gracefully shutdown the server", err)
	}
	log.Info("Gracefully stopped")

}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
