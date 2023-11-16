package config

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type MiddleConf struct {
	Platform  string `json:"platform"`  // 平台类型
	Serve     string `json:"serve"`     // 当前服务
	Root      string `json:"root"`      // 当前服务
	Host      string `json:"host"`      // 主机名
	Port      string `json:"port"`      // 端口号
	DataPort  string `json:"dataPort"`  // 数据库端口
	DataCache string `json:"dataCache"` // 缓存数据库地址
}

func MiddleWare(options MiddleConf) gin.HandlerFunc {
	return func(c *gin.Context) {
		list := reflect.ValueOf(&options)
		elem := list.Elem()
		refType := elem.Type()
		for i := 0; i < refType.NumField(); i++ {
			name := refType.Field(i).Name
			c.Set(name, elem.Field(i).Interface())
		}
		c.Next()
	}
}
