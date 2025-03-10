package file

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/AnnaVyvert/safe-concept-server/internal/http/server/api/response"
	"github.com/AnnaVyvert/safe-concept-server/internal/http/server/middleware"
	"github.com/AnnaVyvert/safe-concept-server/internal/log/sl"
	"github.com/AnnaVyvert/safe-concept-server/internal/storage/file"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// Create user file
// Rreturns error if file exists
func Create(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Update"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Info("bad body id", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		if _, err = fileStorage.Load(file.ID(appID, userID)); err == nil {
			log.Info("file alreay exists")
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		if err = fileStorage.Store(file.ID(appID, userID), data); err != nil {
			log.Info("can not save file", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}
		render.JSON(w, r, resp.FileOK())
	}
}

// Get user file
// Returns error if file does not exists
func Get(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Get"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		var data []byte
		if data, err = fileStorage.Load(file.ID(appID, userID)); err != nil {
			log.Info("file does not exist", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}
		render.JSON(w, r, resp.FileData(data))
	}
}

// Update user file
// Returns error if file does not exists
func Update(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Update"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Info("bad body id", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		var oldData []byte
		if oldData, err = fileStorage.Load(file.ID(appID, userID)); err != nil {
			log.Info("file does not exists", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		if err = fileStorage.Store(file.ID(appID, userID), data); err != nil {
			log.Info("can not save file", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}
		render.JSON(w, r, resp.FileData(oldData))
	}
}

// Delete user file
// Returns error if file does not exists
func Delete(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Get"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		// var oldData []byte read before delete if needed
		if err = fileStorage.Delete(file.ID(appID, userID)); err != nil {
			log.Info("file does not exist", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resp.FileError("invalid request"))
			return
		}
		render.JSON(w, r, resp.FileOK())
	}
}
