package vars

import (
	lib "github.com/AnnaVyvert/safe-concept-server/cmd/server/common_lib"
	"github.com/joho/godotenv"
)

var FsFolderPath string

func Load() error {
	if err := godotenv.Load("../.env"); err != nil {
		return err
	}
	FsFolderPath = lib.GetEnv("FS_FOLDER_PATH")
	return nil
}
