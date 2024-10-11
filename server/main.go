package main

import (
	"browser-remote-server/internal/config"
	"browser-remote-server/internal/http-server/events"
	"browser-remote-server/internal/http-server/handlers/elements/delete"
	"browser-remote-server/internal/http-server/handlers/elements/save"
	"browser-remote-server/internal/http-server/handlers/page"
	"browser-remote-server/internal/http-server/handlers/trigger"
	"browser-remote-server/internal/http-server/middleware"
	"browser-remote-server/internal/storage/jsonstorage"
	"browser-remote-server/lib/logger/handlers/slogpretty"
	"flag"
	"log"
	"log/slog"
	"net/http"
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

	log.Info("server started", slog.String("env", cfg.Env), slog.String("address", cfg.HTTPServer.Address))
	log.Debug("debug logging enabled")

	storage := jsonstorage.New(cfg.StoragePath)
	storage.Init()

	eventController := events.New(log)

	router := mux.NewRouter()

	router.Use(middleware.NewLogger(log))

	router.HandleFunc("/", page.New(log, cfg.HTMLPath))
	router.HandleFunc("/elements/save", save.New(log, storage)).Methods("POST")
	router.HandleFunc("/elements/delete", delete.New(log, storage)).Methods("POST")
	router.HandleFunc("/trigger", trigger.New(log, eventController, storage)).Methods("POST")

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.HTTPServer.Address,
		WriteTimeout: cfg.HTTPServer.Timeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		{
			log = setupPrettySlog()
		}
	case envProd:
		{
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
