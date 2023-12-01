package tools

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type YamlConf struct {
	Namespace string                       `json:"namespace"` // 项目空间名称
	Name      string                       `json:"name"`      // 项目名称
	Version   string                       `json:"version"`   // 版本号
	Root      map[string]string            `json:"root"`      // 根目录信息
	Database  map[string]map[string]string `json:"database"`  // 数据库信息
	Services  map[string]ServeConf         `json:"services"`  // 微服务列表
}

type ServeConf struct {
	Host     string `json:"host"`     // 主域名
	Port     string `json:"port"`     // 服务端口号
	Database string `json:"database"` // 数据库名
	DataPort string `json:"dataPort"` // 数据库端口
}
type mapConf = map[string]map[string]string

func Language(dir string) (mapConf, error) {
	dataBytes, err := os.ReadFile(dir)
	if err != nil {
		return mapConf{}, errors.New("获取语言配置失败")
	}
	options := mapConf{}
	err = yaml.Unmarshal(dataBytes, &options)
	if err != nil {
		return mapConf{}, errors.New("解析语言配置失败")
	}
	return options, err
}

func Yaml(dir string) (YamlConf, error) {
	dataBytes, err := os.ReadFile(dir)
	if err != nil {
		return YamlConf{}, err
	}
	options := YamlConf{}
	err = yaml.Unmarshal(dataBytes, &options)
	if err != nil {
		return YamlConf{}, err
	}
	return options, err
}
