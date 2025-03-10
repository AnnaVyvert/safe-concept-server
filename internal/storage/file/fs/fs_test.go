package fs

import (
	"log/slog"
	"os"
	"path"
	"testing"

	"github.com/AnnaVyvert/safe-concept-server/internal/storage/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ValidFileStorage(t *testing.T, dirName string) file.Storage {
	storagePath := path.Join(os.TempDir(), dirName)
	t.Helper()
	require.NoError(t, os.MkdirAll(storagePath, os.ModePerm))
	storageDir, err := os.MkdirTemp(storagePath, "")
	require.NoError(t, err)
	// TODO(mxd): nop-logger
	return NewFileStorage(slog.Default(), storageDir)
}

// func InvalidFileStorage() TODO(mxd): implimment

func TestStorageLoadAndStore(t *testing.T) {
	testCases := []struct {
		storageType string
		storage     file.Storage
	}{
		{
			storageType: "file",
			storage:     ValidFileStorage(t, "file_test_load_and_store"),
		},
		{
			storageType: "crypted",
			storage:     file.Crypted(ValidFileStorage(t, "crypted_test_load_and_store")),
		},
	}

	assert.Panics(t, func() {
		NewFileStorage(slog.Default(), "/~.../../../")
	})

	for _, tC := range testCases {
		t.Run(tC.storageType, func(t *testing.T) {
			require := require.New(t)

			storage := tC.storage
			expected := []byte("буу ! испугался ?")

			id := file.ID(1, 1)
			_, err := storage.Load(id)
			require.Error(err)

			require.NoError(storage.Store(id, expected))

			data, err := storage.Load(id)
			require.NoError(err)
			require.Equal(expected, data)
		})
	}
}
