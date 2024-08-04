package tools

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type MapConf = map[string]interface{}

func Yaml(dir string) (MapConf, error) {
	dataBytes, err := os.ReadFile(dir)
	if err != nil {
		return MapConf{}, errors.New("query")
	}
	options := MapConf{}
	err = yaml.Unmarshal(dataBytes, &options)
	if err != nil {
		return MapConf{}, errors.New("serialization")
	}
	return options, err
}
