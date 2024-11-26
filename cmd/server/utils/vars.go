package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

var FsFolderPath string = fmt.Sprintf("/var/tmp/www/%s/", "safe-concept-server")

func Load() error {
	if err := godotenv.Load("../.env"); err != nil {
		return err
	}
	FsFolderPath = GetEnv("FS_FOLDER_PATH")
	return nil
}
