package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type YamlConf struct {
	Namespace string                  `json:"namespace"` // 项目空间名称
	Name      string                  `json:"name"`      // 项目名称
	Version   string                  `json:"version"`   // 版本号
	Root_dir  PlatformConf            `json:"root_dir"`  // 根目录信息
	Database  map[string]PlatformConf `json:"database"`  // 数据库信息
	Services  map[string]ServeConf    `json:"services"`  // 微服务列表
}

type ServeConf struct {
	Host     string `json:"host"`     // 主域名
	Port     string `json:"port"`     // 服务端口号
	Database string `json:"database"` // 数据库名
	DataPort string `json:"dataPort"` // 数据库端口
}

type PlatformConf struct {
	Linux   string `json:"linux"`   // linux平台
	Windows string `json:"windows"` // window平台
}

func Yaml(dir string) (interface{}, error) {
	dataBytes, err := os.ReadFile(dir)
	if err != nil {
		return "", err
	}
	options := YamlConf{}
	err = yaml.Unmarshal(dataBytes, &options)
	if err != nil {
		return "", err
	}
	return options, err
}
