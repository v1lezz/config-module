package conf_worker

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

func GetConfig(configPath string) (ServerConfig, error) {
	if validateConfigPath(configPath) {
		return ServerConfig{}, errors.New("incorrect path of config file")
	}
	file, err := os.Open(configPath)
	if err != nil {
		return ServerConfig{}, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	sc := ServerConfig{}
	if err = decoder.Decode(&sc); err != nil {
		return ServerConfig{}, err
	}
	return sc, nil
}

func validateConfigPath(configPath string) bool {
	fileName, err := getNameFile(configPath)
	if err != nil {
		return false
	}
	return validateConfigFileName(fileName)
}

func getNameFile(path string) (string, error) {
	splitted := strings.Split(path, "\\")
	for i := len(splitted) - 1; i >= 0; i-- {
		if splitted[i] != "" {
			return splitted[i], nil
		}
	}
	return "", errors.New("error get config file name")
}

func validateConfigFileName(configFileName string) bool {
	splitted := strings.Split(configFileName, ".")
	if len(splitted) != 2 {
		return false
	}
	return splitted[1] == "yaml" && splitted[0] != ""
}
