package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

var (
	LogEntryCtxKey = &contextKey{"LogEntry"}
)

func WithSlog(entry *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), LogEntryCtxKey, entry)))
		}
		return http.HandlerFunc(fn)
	}
}

func GetSlog(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return nil
	}
	if logger, ok := ctx.Value(LogEntryCtxKey).(*slog.Logger); ok {
		return logger
	}

	return nil
}

func RequestSlog(log *slog.Logger, requestIDKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Info("slog middleware enabled", slog.String("request-id-key", requestIDKey))

		fn := func(w http.ResponseWriter, r *http.Request) {
			withRequestID := middleware.RequestID
			withSlog := WithSlog(log)
			withLogger := middleware.RequestLogger(newLogFormatter(log, requestIDKey))
			withRequestID(
				withSlog(
					withLogger(
						next,
					),
				),
			).ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

type slogEntry struct {
	logger *slog.Logger
	debug  bool
}

func (s *slogEntry) Write(status int, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	type Speed string

	const (
		fast     Speed = "fast"
		long     Speed = "long"
		veryLong Speed = "very long"
	)

	var speed Speed
	if elapsed < 500*time.Millisecond {
		speed = fast
	} else if elapsed < 5*time.Second {
		speed = long
	} else {
		speed = veryLong
	}

	entry := s.logger.With(
		slog.Int("status", status),
		slog.Duration("elapsed", elapsed),
		slog.String("speed", string(speed)),
	)

	if s.debug {
		entry = entry.With(slog.Int("bytes", bytes))
	}

	switch {
	case status < 200:
		entry.Info("request completed")
	case status < 300:
		entry.Info("request completed")
	case status < 400 || speed == long:
		entry.Warn("request completed")
	case status < 500 || speed == veryLong:
		entry.Error("request completed")
	default:
		entry.Info("request completed")
	}
}

func (s *slogEntry) Panic(rvr interface{}, stack []byte) {
	s.logger.Error(
		"panic",
		slog.Any("recover", rvr),
		slog.String("stack", string(debug.Stack())),
	)
}

var _ middleware.LogEntry = new(slogEntry)

type logFormatter struct {
	logger       *slog.Logger
	requestIDKey string
}

func (l *logFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	entry := l.logger.With(
		slog.String("method", r.Method),
		slog.String("url", fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)),
		slog.String("remote_addr", r.RemoteAddr),
		slog.String(l.requestIDKey, middleware.GetReqID(r.Context())),
	)

	entry.Info("requested")

	return &slogEntry{logger: entry, debug: entry.Enabled(r.Context(), slog.LevelDebug)}
}

var _ middleware.LogFormatter = new(logFormatter)

func newLogFormatter(log *slog.Logger, requestIDKey string) middleware.LogFormatter {
	log.Debug("new log formatter", slog.String("request-id-key", requestIDKey))
	return &logFormatter{
		logger:       log,
		requestIDKey: requestIDKey,
	}
}
