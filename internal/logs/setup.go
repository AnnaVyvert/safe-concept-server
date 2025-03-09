package logs

import (
	"log/slog"
	"os"

	"github.com/AnnaVyvert/safe-concept-server/internal/config"
)

func Setup(cfg *config.Config) (log *slog.Logger) {
	switch cfg.Env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
