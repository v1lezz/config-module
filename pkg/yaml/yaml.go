package yaml

import (
	"errors"
	"github.com/v1lezz/config-module/internal/validate"
	"github.com/v1lezz/config-module/pkg/domain/models"
	"gopkg.in/yaml.v3"
	"os"
)

func GetConfig(configPath string) (models.ServerConfig, error) {
	if validate.ValidateConfigPath(configPath) {
		return models.ServerConfig{}, errors.New("incorrect path config file")
	}
	file, err := os.Open(configPath)
	if err != nil {
		return models.ServerConfig{}, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	sc := models.ServerConfig{}
	if err = decoder.Decode(&sc); err != nil {
		return models.ServerConfig{}, err
	}
	return sc, nil
}
