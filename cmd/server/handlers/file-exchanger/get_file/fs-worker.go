package get_file

import (
	"os"

	lib "github.com/AnnaVyvert/safe-concept-server/cmd/server/common_lib"
)

func getFileFromFS(fileName string) []byte {
	fsFolderPath := lib.GetEnv("FS_FOLDER_PATH")

	file, err := os.ReadFile(fsFolderPath + fileName)

	lib.PanicIfError(err)

	return file
}
