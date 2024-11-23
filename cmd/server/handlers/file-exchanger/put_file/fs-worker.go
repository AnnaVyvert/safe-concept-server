package put_file

import (
	"os"

	lib "github.com/AnnaVyvert/safe-concept-server/cmd/server/common_lib"
)

func writeFileOnPath(fsFolderPath string, fileName string, content string) {
	err := os.MkdirAll(fsFolderPath, 0700)
	lib.PanicIfError(err)

	f, err := os.Create(fsFolderPath + fileName)
	lib.PanicIfError(err)

	err = os.WriteFile(fsFolderPath+fileName, []byte(content), 0600)
	lib.PanicIfError(err)

	f.Close()
}

func writeFileOnFS(fileNameToSave string, fileContentToSave string) {
	fsFolderPath := lib.GetEnv("FS_FOLDER_PATH")

	writeFileOnPath(fsFolderPath, fileNameToSave, fileContentToSave)
}
