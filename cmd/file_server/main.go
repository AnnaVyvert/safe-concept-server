package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/AnnaVyvert/safe-concept-server/internal/config"
	"github.com/AnnaVyvert/safe-concept-server/internal/http/server/handlers/file"
	"github.com/AnnaVyvert/safe-concept-server/internal/http/server/middleware"
	"github.com/AnnaVyvert/safe-concept-server/internal/log/sl"
	fs_storage "github.com/AnnaVyvert/safe-concept-server/internal/storage/file/fs"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/render"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// config
	var cfg *config.Config = config.MustLoad()

	FsFolderPath := path.Join("/var/tmp/www/", "safe-concept-server")

	// log
	// FIXME(mxd)
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// storage
	fileStorage := fs_storage.NewFileStorage(log.With(slog.String("module", "storage")), FsFolderPath)

	// handlers
	router := chi.NewRouter()
	if cfg.Env == "dev" {
		slog.SetDefault(log)
		router.Use(middleware.WithSlog(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))))
		router.Use(chi_middleware.Logger)
	} else {
		router.Use(middleware.RequestSlog(log, cfg.Log.RequestIDKey))
	}
	router.Use(chi_middleware.Recoverer)

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
