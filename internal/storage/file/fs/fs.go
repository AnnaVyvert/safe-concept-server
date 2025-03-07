package fs

import (
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/AnnaVyvert/safe-concept-server/internal/log/sl"
	"github.com/AnnaVyvert/safe-concept-server/internal/storage/file"
)

const perm = 0754

func fileDir(id file.FileID) string {
	return fmt.Sprintf("%d", id.App)
}
func filePath(id file.FileID) string {
	return fmt.Sprintf("%s/%d", fileDir(id), id.User)
}

var _ file.Storage = fileStorage{}

type fileStorage struct {
	log        *slog.Logger
	storageDir string
}

func NewFileStorage(log *slog.Logger, storageDir string) file.Storage {
	const op = "storage.file.fs.NewFileStorage"
	if err := os.MkdirAll(storageDir, perm); err != nil {
		panic(fmt.Sprintf(
			"can not make dir %q: %q",
			storageDir,
			fmt.Errorf("%s: %w", op, err),
		))
	}
	return fileStorage{
		log: log.With(
			slog.String("storage_type", "fs"),
			slog.String("storage_path", storageDir),
		),
		storageDir: storageDir,
	}
}

func (f fileStorage) Load(id file.FileID) (file.Value, error) {
	const op = "storage.file.fs.Load"
	filePath := path.Join(f.storageDir, filePath(id))
	data, err := os.ReadFile(filePath)
	if err != nil {
		f.log.Warn("can not read file",
			slog.String("op", op),
			slog.String("path", filePath),
			sl.Err(err),
		)
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}

func (f fileStorage) Store(id file.FileID, data file.Value) error {
	const op = "storage.file.fs.Store"
	if err := os.MkdirAll(path.Join(f.storageDir, fileDir(id)), perm); err != nil {
		f.log.Debug("can not create file dir", sl.Err(err))
	}
	filePath := path.Join(f.storageDir, filePath(id))
	err := os.WriteFile(filePath, data, perm)
	if err != nil {
		f.log.Warn("can not write file",
			slog.String("op", op),
			slog.String("path", filePath),
			sl.Err(err),
		)
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (f fileStorage) Delete(id file.FileID) error {
	const op = "storage.file.fs.Delete"
	filePath := path.Join(f.storageDir, filePath(id))
	err := os.Remove(filePath)
	if err != nil {
		f.log.Warn("can not remove file",
			slog.String("op", op),
			slog.String("path", filePath),
			sl.Err(err),
		)
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
