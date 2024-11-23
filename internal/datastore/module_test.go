package datastore

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

var _ Identifier = new(Id)

type Id string

func (id Id) Identity() string {
	return string(id)
}

func FileStorage(t *testing.T, dirName string) Storage {
	storagePath := path.Join(os.TempDir(), dirName)
	t.Helper()
	require.NoError(t, os.MkdirAll(storagePath, os.ModePerm))
	storageDir, err := os.MkdirTemp(storagePath, "")
	require.NoError(t, err)
	return NewFileStorage(storageDir)
}

func TestStorageLoadAndStore(t *testing.T) {
	testCases := []struct {
		storageType string
		storage     Storage
	}{
		{
			storageType: "file",
			storage:     FileStorage(t, "file_test_load_and_store"),
		},
		{
			storageType: "crypted",
			storage:     Crypted(FileStorage(t, "crypted_test_load_and_store")),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.storageType, func(t *testing.T) {
			require := require.New(t)

			storage := tC.storage
			expected := []byte("буу ! испугался ?")

			id := Id("foo")
			_, err := storage.Load(id)
			require.Error(err)

			require.NoError(storage.Store(id, expected))

			data, err := storage.Load(id)
			require.NoError(err)
			require.Equal(expected, data)
		})
	}
}
