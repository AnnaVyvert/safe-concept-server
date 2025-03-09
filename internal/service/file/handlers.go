package file

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AnnaVyvert/safe-concept-server/internal/http/server/middleware"
	"github.com/AnnaVyvert/safe-concept-server/internal/log/sl"
	"github.com/AnnaVyvert/safe-concept-server/internal/storage/file"
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type jsonValue map[string]any

func parseValue(reader io.Reader, tee io.Writer) (jsonValue, error) {
	data := make(jsonValue)

	if tee != nil {
		reader = io.TeeReader(reader, tee)
	}
	err := json.NewDecoder(reader).Decode(&data)
	return data, err
}

// Create user file
// Rreturns error if file exists
func Create(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Create"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		var dataBytes bytes.Buffer

		if _, err = parseValue(r.Body, &dataBytes); err != nil {
			log.Info("invalid body json", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		if _, err = fileStorage.Load(file.ID(appID, userID)); err == nil {
			log.Info("file alreay exists")
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		if err = fileStorage.Store(file.ID(appID, userID), dataBytes.Bytes()); err != nil {
			log.Info("can not save file", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, FileError("invalid request"))
			return
		}
		render.JSON(w, r, FileOK())
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
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		var dataBytes []byte
		if dataBytes, err = fileStorage.Load(file.ID(appID, userID)); err != nil {
			log.Info("file does not exist", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		var data jsonValue
		if data, err = parseValue(bytes.NewReader(dataBytes), nil); err != nil {
			panic("unreachable")
		}

		render.JSON(w, r, FileData(data))
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
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		dataBytes := bytes.NewBuffer(nil)
		if _, err := parseValue(io.TeeReader(r.Body, dataBytes), nil); err != nil {
			log.Info("invalid body json", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		if _, err = fileStorage.Load(file.ID(appID, userID)); err != nil {
			log.Info("file does not exists", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		if err = fileStorage.Store(file.ID(appID, userID), dataBytes.Bytes()); err != nil {
			log.Info("can not save file", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		render.JSON(w, r, FileOK())
	}
}

// Delete user file
// Returns error if file does not exists
func Delete(fileStorage file.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.file.Delete"
		log := middleware.GetSlog(r.Context()).With(
			slog.String("op", op),
			slog.String("request_id", chi_middleware.GetReqID(r.Context())),
		)

		var appID int
		var err error
		if appID, err = strconv.Atoi(chi.URLParam(r, "app_id")); err != nil {
			log.Info("bad app id", sl.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, FileError("invalid request"))
			return
		}

		userID := 0 // TODO(mxd): Auth middleware

		// var oldData []byte read before delete if needed
		if err = fileStorage.Delete(file.ID(appID, userID)); err != nil {
			log.Info("file does not exist", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, FileError("invalid request"))
			return
		}
		render.JSON(w, r, FileOK())
	}
}
