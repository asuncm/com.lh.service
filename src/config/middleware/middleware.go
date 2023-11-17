package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		host := c.Request.Host
		c.Header("Access-Control-Allow-Origin", host)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, SSID, Verify, UUID")
		c.Header("Access-Control-Allow-Methods", method)
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
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
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
