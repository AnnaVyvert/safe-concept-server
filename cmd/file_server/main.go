package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/AnnaVyvert/safe-concept-server/internal/config"
	"github.com/AnnaVyvert/safe-concept-server/internal/http/server/middleware"
	"github.com/AnnaVyvert/safe-concept-server/internal/logs"
	"github.com/AnnaVyvert/safe-concept-server/internal/logs/sl"
	"github.com/AnnaVyvert/safe-concept-server/internal/service/file"
	"github.com/AnnaVyvert/safe-concept-server/internal/storage/file/fs"
)

func main() {
	var cfg *config.Config = config.MustLoad()

	log := logs.Setup(cfg)

	// storage
	fileStorage := fs.NewFileStorage(log.With(slog.String("module", "storage")), cfg.FsFolderPath)

	// handlers
	router := chi.NewRouter()
	if cfg.Env == "local" {
		slog.SetDefault(log)
		router.Use(chi_middleware.RequestID)
		router.Use(middleware.WithSlog(log))
		router.Use(chi_middleware.Logger)
	} else {
		router.Use(middleware.RequestSlog(log, cfg.Log.RequestIDKey))
	}
	router.Use(chi_middleware.Recoverer)

	// TODO(mxd): Можно хранить app_id в JWT токене
	router.Route("/app/{app_id}/file", func(r chi.Router) {
		r.Post("/", file.Create(fileStorage))
		r.Get("/", file.Get(fileStorage))
		r.Put("/", file.Update(fileStorage))
		r.Delete("/", file.Delete(fileStorage))
	})

	// server
	log.Info("starting server", slog.String("address", cfg.Address), slog.String("env", cfg.Env))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen and serve have stopped", sl.Err(err))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}

	log.Info("server stopped")
}
