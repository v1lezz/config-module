package app

import (
	"errors"
	"strings"
)

func GetNameFile(path string) (string, error) {
	splitted := strings.Split(path, "\\")
	for i := len(splitted) - 1; i >= 0; i-- {
		if splitted[i] != "" {
			return splitted[i], nil
		}
	}
	return "", errors.New("error get config file name")
}
