package main

import (
	"browser-remote-server/internal/config"
	"browser-remote-server/internal/http-server/handlers/elements/save"
	"browser-remote-server/internal/storage/jsonstorage"
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/gorilla/mux"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	flag.Parse()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("config error: ", err)
	}

	log := setupLogger(cfg.Env)

	log.Info("server started", slog.String("env", cfg.Env))
	log.Debug("debug logging enabled")

	storage := jsonstorage.New(cfg.StoragePath)

	router := mux.NewRouter()

	router.HandleFunc("save", save.New(log, storage))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		{
			log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}
	case envProd:
		{
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}
	}

	return log
}
