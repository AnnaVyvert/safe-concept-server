package put_file

import (
	"os"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/utils"
)

func writeFileOnPath(fsFolderPath string, fileName string, content string) {
	err := os.MkdirAll(fsFolderPath, 0700)
	utils.PanicIfError(err)

	f, err := os.Create(fsFolderPath + fileName)
	utils.PanicIfError(err)

	err = os.WriteFile(fsFolderPath+fileName, []byte(content), 0600)
	utils.PanicIfError(err)

	f.Close()
}

func writeFileOnFS(fileNameToSave string, fileContentToSave string) {
	writeFileOnPath(utils.FsFolderPath, fileNameToSave, fileContentToSave)
}
