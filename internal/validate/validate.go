package validate

import (
	"github.com/v1lezz/config-module/internal/app"
	"strings"
)

func ValidateConfigPath(configPath string) bool {
	fileName, err := app.GetNameFile(configPath)
	if err != nil {
		return false
	}
	return validateConfigFileName(fileName)
}

func validateConfigFileName(configFileName string) bool {
	splitted := strings.Split(configFileName, ".")
	if len(splitted) != 2 {
		return false
	}
	return splitted[1] == "yaml" && splitted[0] != ""
}
