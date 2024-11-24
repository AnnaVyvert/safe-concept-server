package get_file

import (
	"os"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/utils"
)

func getFileFromFS(fileName string) []byte {
	file, err := os.ReadFile(utils.FsFolderPath + fileName)

	utils.PanicIfError(err)

	return file
}
